package middleware

import (
	"net/http"
	"strings"

	"GoCommerce/backend/utils"
)

// ValidateToken middleware para validar JWT
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Token faltante", nil)
			return
		}

		parts := strings.Split(tokenHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Formato de token inválido", nil)
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Token inválido o expirado", nil)
			return
		}

		ctx := utils.WithUserClaims(r.Context(), claims)
		next(w, r.WithContext(ctx))
	}
}
