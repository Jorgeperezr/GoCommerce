package middleware

import (
	"GoCommerce/backend/utils"
	"log"
	"net/http"
	"runtime/debug"
)

// ErrorHandler maneja errores globales de la aplicación
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error inesperado: %v\nStack trace: %s", err, debug.Stack())
				utils.ErrorResponse(w, http.StatusInternalServerError, "Error interno del servidor", nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
