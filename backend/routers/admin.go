package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterAdminRoutes registra las rutas de administración.
// El router recibido ya está montado bajo /admin desde main.go.
// Todas requieren token válido y rol de administrador.
func RegisterAdminRoutes(router *mux.Router) {
	// Panel de control
	router.HandleFunc("/dashboard", middleware.ValidateToken(middleware.AdminOnly(controllers.Dashboard))).Methods("GET")

	// Gestión de usuarios
	router.HandleFunc("/users/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.DeleteUser))).Methods("DELETE")

	// Transacciones
	router.HandleFunc("/transactions", middleware.ValidateToken(middleware.AdminOnly(controllers.GetAllTransactions))).Methods("GET")
	router.HandleFunc("/transactions/{user_id:[0-9]+}", middleware.ValidateToken(middleware.AdminOnly(controllers.GetTransactionHistory))).Methods("GET")
}
