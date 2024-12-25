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

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func startDatabase() {
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connected successfully")
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Middlewares globales
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ErrorHandler)

	// Registrar rutas
	routers.RegisterAuthRoutes(router)
	routers.RegisterProductRoutes(router)

	// Ruta de prueba
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ruta de prueba funcionando"))
	}).Methods("GET")

	return router
}

func startServer(router *mux.Router) {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if _, err := fmt.Sscanf(port, "%d", new(int)); err != nil {
		log.Fatalf("Puerto APP_PORT inválido: %v", err)
	}

	log.Printf("Servidor corriendo en http://localhost:%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func main() {
	loadEnv()
	startDatabase()
	router := setupRouter()
	startServer(router)
}
