package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes registra las rutas relacionadas con productos.
// El router recibido ya está montado bajo /products desde main.go.
func RegisterProductRoutes(router *mux.Router) {
	// Rutas públicas
	router.HandleFunc("", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.GetProductByID).Methods("GET")

	// Rutas protegidas para administradores
	router.HandleFunc("", middleware.ValidateToken(middleware.AdminOnly(controllers.CreateProduct))).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.UpdateProduct))).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.DeleteProduct))).Methods("DELETE")
}
