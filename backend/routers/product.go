package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes registra las rutas relacionadas con productos
func RegisterProductRoutes(router *mux.Router) {
	productRouter := router.PathPrefix("/products").Subrouter()

	// Rutas públicas
	productRouter.HandleFunc("", controllers.GetProducts).Methods("GET")
	productRouter.HandleFunc("/{id}", controllers.GetProductByID).Methods("GET")

	// Rutas protegidas para administradores
	productRouter.HandleFunc("", middleware.ValidateToken(middleware.AdminOnly(controllers.CreateProduct))).Methods("POST")
	productRouter.HandleFunc("/{id}", middleware.ValidateToken(middleware.AdminOnly(controllers.UpdateProduct))).Methods("PUT")
	productRouter.HandleFunc("/{id}", middleware.ValidateToken(middleware.AdminOnly(controllers.DeleteProduct))).Methods("DELETE")
}
