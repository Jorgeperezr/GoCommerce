package api

import (
	"GoCommerce/controllers"
	"net/http"
)

// ConfigurarRutas configura todas las rutas de la aplicación
func ConfigurarRutas() {
	http.HandleFunc("/productos", controllers.ObtenerProductos)
	http.HandleFunc("/productos/eliminar", controllers.EliminarProducto)
	http.HandleFunc("/historial", controllers.ObtenerHistorial)
}
