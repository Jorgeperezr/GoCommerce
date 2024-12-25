package middleware

import (
	"GoCommerce/backend/utils"
	"net/http"
)

// AdminOnly restringe el acceso a usuarios con rol de administrador
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := utils.GetUserClaims(r)
		if !ok || claims == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "No autorizado", nil)
			return
		}

		if claims.Role != "admin" {
			utils.ErrorResponse(w, http.StatusForbidden, "Acceso restringido a administradores", nil)
			return
		}

		next(w, r)
	}
}

// UserOnly restringe el acceso a usuarios autenticados
func UserOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := utils.GetUserClaims(r)
		if !ok || claims == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "No autorizado", nil)
			return
		}

		next(w, r)
	}
}
