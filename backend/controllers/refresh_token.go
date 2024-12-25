package controllers

import (
	"GoCommerce/backend/utils"
	"net/http"
)

// RefreshToken maneja la renovación de tokens JWT
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetUserClaims(r)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Token inválido o expirado", nil)
		return
	}

	// Generar un nuevo token
	token, err := utils.GenerateToken(claims.Username, claims.Role)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al generar un nuevo token", nil)
		return
	}

	utils.SuccessResponse(w, "Token renovado correctamente", map[string]string{
		"token": token,
	})
}
