package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterOrderRoutes registra las rutas relacionadas con pedidos.
// El router recibido ya está montado bajo /orders desde main.go.
func RegisterOrderRoutes(router *mux.Router) {
	// Cualquier usuario autenticado puede crear y consultar pedidos
	router.HandleFunc("", middleware.ValidateToken(middleware.UserOnly(controllers.CreateOrder))).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", middleware.ValidateToken(middleware.UserOnly(controllers.GetOrder))).Methods("GET")

	// Solo administradores pueden listar, actualizar y eliminar pedidos
	router.HandleFunc("", middleware.ValidateToken(middleware.AdminOnly(controllers.GetAllOrders))).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.UpdateOrder))).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.DeleteOrder))).Methods("DELETE")
}
