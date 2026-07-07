package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"GoCommerce/backend/database"
	"GoCommerce/backend/migrations"
	"GoCommerce/backend/routers"

	"github.com/gorilla/mux"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	os.Setenv("SECRET_KEY", "clave_de_pruebas_no_usar_en_produccion")
	os.Setenv("DB_DIALECT", "sqlite")
	os.Setenv("DB_NAME", "file::memory:?cache=shared")

	if err := database.ConnectDB(); err != nil {
		fmt.Println("error conectando la base de datos de pruebas:", err)
		os.Exit(1)
	}
	migrations.InitMigration()

	router := mux.NewRouter()
	routers.RegisterAuthRoutes(router.PathPrefix("/auth").Subrouter())
	routers.RegisterProductRoutes(router.PathPrefix("/products").Subrouter())
	routers.RegisterOrderRoutes(router.PathPrefix("/orders").Subrouter())
	routers.RegisterAdminRoutes(router.PathPrefix("/admin").Subrouter())

	server = httptest.NewServer(router)
	code := m.Run()
	server.Close()
	os.Exit(code)
}

func postJSON(t *testing.T, path, token string, body any) *http.Response {
	t.Helper()
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", server.URL+path, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("POST %s: %v", path, err)
	}
	return resp
}

func getJSON(t *testing.T, path, token string) *http.Response {
	t.Helper()
	req, _ := http.NewRequest("GET", server.URL+path, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET %s: %v", path, err)
	}
	return resp
}

func login(t *testing.T, username, password string) string {
	t.Helper()
	resp := postJSON(t, "/auth/login", "", map[string]string{
		"username": username,
		"password": password,
	})
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login de %s falló con estado %d", username, resp.StatusCode)
	}
	var body struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&body)
	if body.Token == "" {
		t.Fatal("login no devolvió token")
	}
	return body.Token
}

// El usuario admin sembrado puede iniciar sesión y ver su perfil
func TestLoginAndProfile(t *testing.T) {
	token := login(t, "admin", "Admin1234")

	resp := getJSON(t, "/auth/profile", token)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("perfil devolvió estado %d", resp.StatusCode)
	}
	var body struct {
		Data struct {
			Username string `json:"username"`
			Role     string `json:"role"`
		} `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&body)
	if body.Data.Username != "admin" || body.Data.Role != "admin" {
		t.Fatalf("perfil inesperado: %+v", body.Data)
	}
}

// Un usuario registrado puede iniciar sesión inmediatamente (sin doble hash)
func TestRegisterThenLogin(t *testing.T) {
	resp := postJSON(t, "/auth/register", "", map[string]string{
		"username": "cliente1",
		"password": "Cliente1234",
	})
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("registro falló con estado %d", resp.StatusCode)
	}
	login(t, "cliente1", "Cliente1234")
}

// El registro nunca expone la contraseña ni su hash
func TestRegisterDoesNotLeakPassword(t *testing.T) {
	resp := postJSON(t, "/auth/register", "", map[string]string{
		"username": "cliente2",
		"password": "Cliente1234",
	})
	defer resp.Body.Close()
	var raw map[string]any
	json.NewDecoder(resp.Body).Decode(&raw)
	data, _ := raw["data"].(map[string]any)
	if _, exists := data["Password"]; exists {
		t.Fatal("la respuesta de registro expone el campo Password")
	}
	if _, exists := data["password"]; exists {
		t.Fatal("la respuesta de registro expone el campo password")
	}
}

// Solo un administrador puede crear productos; un pedido descuenta stock
func TestProductAndOrderFlow(t *testing.T) {
	adminToken := login(t, "admin", "Admin1234")

	// Crear producto como admin
	resp := postJSON(t, "/products", adminToken, map[string]any{
		"name":  "Teclado mecánico",
		"price": 49.99,
		"stock": 10,
	})
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("crear producto devolvió estado %d", resp.StatusCode)
	}
	var product struct {
		ID    uint    `json:"id"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}
	json.NewDecoder(resp.Body).Decode(&product)
	resp.Body.Close()
	if product.ID == 0 {
		t.Fatal("el producto creado no tiene id")
	}

	// Un usuario normal no puede crear productos
	userToken := login(t, "empleado1", "Empleado1234")
	resp = postJSON(t, "/products", userToken, map[string]any{
		"name":  "Producto pirata",
		"price": 1.0,
	})
	resp.Body.Close()
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("un usuario sin rol admin creó un producto (estado %d)", resp.StatusCode)
	}

	// Crear pedido como usuario: el servidor calcula el total y descuenta stock
	resp = postJSON(t, "/orders", userToken, map[string]any{
		"product_id": product.ID,
		"quantity":   3,
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("crear pedido devolvió estado %d", resp.StatusCode)
	}
	var orderBody struct {
		Data struct {
			TotalPrice float64 `json:"total_price"`
		} `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&orderBody)
	resp.Body.Close()
	want := 49.99 * 3
	if orderBody.Data.TotalPrice != want {
		t.Fatalf("total_price = %v, esperado %v", orderBody.Data.TotalPrice, want)
	}

	// El stock quedó descontado
	resp = getJSON(t, fmt.Sprintf("/products/%d", product.ID), "")
	json.NewDecoder(resp.Body).Decode(&product)
	resp.Body.Close()
	if product.Stock != 7 {
		t.Fatalf("stock = %d, esperado 7", product.Stock)
	}

	// Un pedido que excede el stock es rechazado
	resp = postJSON(t, "/orders", userToken, map[string]any{
		"product_id": product.ID,
		"quantity":   999,
	})
	resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("pedido sin stock suficiente devolvió estado %d", resp.StatusCode)
	}
}

// Las rutas protegidas rechazan peticiones sin token
func TestProtectedRoutesRequireToken(t *testing.T) {
	for _, path := range []string{"/auth/profile", "/admin/dashboard"} {
		resp := getJSON(t, path, "")
		resp.Body.Close()
		if resp.StatusCode != http.StatusUnauthorized {
			t.Fatalf("%s sin token devolvió estado %d, esperado 401", path, resp.StatusCode)
		}
	}
}
