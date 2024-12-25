package controllers

import (
	"GoCommerce/backend/utils"
	"net/http"
)

// ProfileHandler maneja las solicitudes para obtener el perfil del usuario autenticado
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetUserClaims(r)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "No autorizado", nil)
		return
	}

	utils.SuccessResponse(w, "Perfil obtenido correctamente", map[string]string{
		"username": claims.Username,
		"role":     claims.Role,
	})
}

// GetProfile obtiene la información del perfil del usuario autenticado
func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetUserClaims(r)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Token inválido o expirado", nil)
		return
	}

	profile := map[string]interface{}{
		"username": claims.Username,
		"role":     claims.Role,
	}

	utils.SuccessResponse(w, "Perfil obtenido exitosamente", profile)
}
