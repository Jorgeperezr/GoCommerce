package api

import (
	"GoCommerce/controllers"
	"fmt"
	"net/http"
)

// ConfigurarRutas define todas las rutas de la API
func ConfigurarRutas() {
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/productos", controllers.ObtenerProductosHandler)
	http.HandleFunc("/historial", controllers.HistorialHandler)

	// Ruta predeterminada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bienvenido a GoCommerce API")
	})

	// Ruta para manejar rutas no encontradas
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Página no encontrada", http.StatusNotFound)
	})
}
