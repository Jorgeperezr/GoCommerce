package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"github.com/gorilla/mux"
)

// UpdateUser actualiza la información de un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Usuario no encontrado", err.Error())
		return
	}

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	if updatedData.Username != "" {
		user.Username = updatedData.Username
	}
	if updatedData.Password != "" {
		user.Password = updatedData.Password
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al actualizar el usuario", err.Error())
		return
	}

	utils.SuccessResponse(w, "Usuario actualizado exitosamente", user)
}
