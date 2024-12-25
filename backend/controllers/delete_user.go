package controllers

import (
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"github.com/gorilla/mux"
)

// DeleteUser elimina un usuario por su ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Usuario no encontrado", err.Error())
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al eliminar el usuario", err.Error())
		return
	}

	utils.SuccessResponse(w, "Usuario eliminado exitosamente", nil)
}
