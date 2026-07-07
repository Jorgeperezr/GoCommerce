package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"GoCommerce/backend/database"
	"GoCommerce/backend/middleware"
	"GoCommerce/backend/migrations"
	"GoCommerce/backend/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// loadEnv carga las variables de entorno desde el archivo .env
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	if os.Getenv("SECRET_KEY") == "" {
		log.Fatal("SECRET_KEY no está configurada: defínela en .env o en el entorno")
	}
}

// startDatabase conecta la base de datos y crea los datos iniciales
func startDatabase() {
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connected successfully")

	migrations.InitMigration()
}

// frontendDir devuelve el directorio de archivos estáticos del frontend
func frontendDir() string {
	if dir := os.Getenv("FRONTEND_DIR"); dir != "" {
		return dir
	}
	return "../frontend"
}

// setupRouter configura las rutas y middlewares
func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Middlewares globales
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ErrorHandler)

	frontend := frontendDir()

	// Recursos estáticos (CSS y JS)
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(frontend, "css")))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(filepath.Join(frontend, "js")))))

	// Páginas del frontend
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontend, "index.html"))
	}).Methods("GET")
	router.HandleFunc("/login.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontend, "auth", "login.html"))
	}).Methods("GET")
	router.HandleFunc("/admin.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontend, "admin.html"))
	}).Methods("GET")

	// Ruta de prueba
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ruta de prueba funcionando correctamente"))
	}).Methods("GET")

	// Rutas de la API
	routers.RegisterAuthRoutes(router.PathPrefix("/auth").Subrouter())
	routers.RegisterProductRoutes(router.PathPrefix("/products").Subrouter())
	routers.RegisterOrderRoutes(router.PathPrefix("/orders").Subrouter())
	routers.RegisterAdminRoutes(router.PathPrefix("/admin").Subrouter())

	// Ruta para manejar 404
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Ruta no encontrada: %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Ruta no encontrada"))
	})

	return router
}

// startServer inicia el servidor HTTP
func startServer(router *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}

	serverAddress := fmt.Sprintf(":%s", port)
	log.Printf("Servidor corriendo en http://localhost%s", serverAddress)

	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// main es el punto de entrada principal de la aplicación
func main() {
	// Cargar variables de entorno
	loadEnv()

	// Iniciar conexión a la base de datos y datos iniciales
	startDatabase()

	// Configurar y lanzar el servidor
	router := setupRouter()
	startServer(router)
}
