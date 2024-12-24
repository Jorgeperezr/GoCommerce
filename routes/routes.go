package routes

import (
	"net/http"

	"github.com/jorge/GoCommerce/controllers"
)

func ConfigurarRutas() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/productos", controllers.ObtenerProductosHandler)
	http.HandleFunc("/historial", controllers.HistorialHandler)
	// Add more routes as needed

	// Ruta predeterminada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	// Ruta para manejar rutas no encontradas
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Página no encontrada", http.StatusNotFound)
	})
}
