package routers

import (
	"GoCommerce/backend/controllers"
	"GoCommerce/backend/middleware"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes registra las rutas relacionadas con autenticación.
// El router recibido ya está montado bajo /auth desde main.go.
func RegisterAuthRoutes(router *mux.Router) {
	// Rutas públicas
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")

	// Rutas protegidas: requieren un token válido
	router.HandleFunc("/refresh-token", middleware.ValidateToken(controllers.RefreshToken)).Methods("POST")
	router.HandleFunc("/profile", middleware.ValidateToken(controllers.GetProfile)).Methods("GET")
}
