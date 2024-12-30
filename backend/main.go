package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"GoCommerce/backend/database"
	"GoCommerce/backend/middleware"
	"GoCommerce/backend/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// loadEnv carga las variables de entorno desde el archivo .env
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

// startDatabase conecta la base de datos
func startDatabase() {
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connected successfully")
}

// setupRouter configura las rutas y middlewares
func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Middlewares globales
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ErrorHandler)

	// ✅ Servir archivos estáticos desde /frontend
	fs := http.FileServer(http.Dir("../frontend"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// ✅ Mapear recursos estáticos (CSS, JS, iconos)
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("../frontend/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("../frontend/js"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("../frontend/images"))))
	router.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir("../frontend/favicon.ico")))

	// ✅ Ruta raíz que carga index.html
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ruta raíz '/' llamada")
		http.ServeFile(w, r, "../frontend/index.html")
	}).Methods("GET")

	// ✅ Ruta de prueba
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ruta '/test' llamada")
		w.Write([]byte("Ruta de prueba funcionando correctamente"))
	}).Methods("GET")

	// ✅ Rutas de autenticación
	log.Println("Registrando rutas de autenticación")
	authRouter := router.PathPrefix("/auth").Subrouter()
	routers.RegisterAuthRoutes(authRouter)

	// ✅ Rutas de productos
	log.Println("Registrando rutas de productos")
	productRouter := router.PathPrefix("/products").Subrouter()
	routers.RegisterProductRoutes(productRouter)

	// ✅ Ruta para manejar 404
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

	// Iniciar conexión a la base de datos
	startDatabase()

	// Configurar y lanzar el servidor
	router := setupRouter()
	startServer(router)
}
