package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes registra las rutas relacionadas con autenticación
func RegisterAuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	// Rutas públicas
	authRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	authRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	authRouter.HandleFunc("/refresh-token", controllers.RefreshToken).Methods("POST")
	authRouter.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Rutas protegidas
	authRouter.HandleFunc("/profile", middleware.ValidateToken(controllers.GetProfile)).Methods("GET")
}
