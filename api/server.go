package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Middleware para el registro de solicitudes
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// Middleware para manejar CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// IniciarServidor inicializa el servidor HTTP
func IniciarServidor() {
	// Cargar rutas configuradas
	ConfigurarRutas()

	// Definir servidor HTTP
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      corsMiddleware(loggerMiddleware(http.DefaultServeMux)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Servidor iniciado en http://localhost:%s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}
