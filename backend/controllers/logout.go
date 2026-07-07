package controllers

import (
	"GoCommerce/backend/utils"
	"net/http"
)

// Logout permite cerrar la sesión del usuario
func Logout(w http.ResponseWriter, r *http.Request) {
	// Para un sistema basado en JWT, el cierre de sesión es manejado en el cliente.
	// Sin embargo, podemos sugerir al cliente invalidar el token borrándolo del almacenamiento local.

	w.Header().Set("Authorization", "")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	utils.SuccessResponse(w, "Sesión cerrada exitosamente", nil)
}
