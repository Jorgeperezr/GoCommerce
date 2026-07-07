package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"github.com/gorilla/mux"
)

// UpdateUserRequest datos admitidos al actualizar un usuario
type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// UpdateUser actualiza la información de un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Usuario no encontrado", err.Error())
		return
	}

	var updatedData UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	if updatedData.Username != "" {
		if err := utils.ValidateUsername(updatedData.Username); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Nombre de usuario inválido", err.Error())
			return
		}
		user.Username = updatedData.Username
	}
	if updatedData.Password != "" {
		if err := utils.ValidatePassword(updatedData.Password); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Contraseña inválida", err.Error())
			return
		}
		if err := user.SetPassword(updatedData.Password); err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error al cifrar la contraseña", err.Error())
			return
		}
	}
	if updatedData.Role == utils.RoleAdmin || updatedData.Role == utils.RoleUser {
		user.Role = updatedData.Role
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al actualizar el usuario", err.Error())
		return
	}

	utils.SuccessResponse(w, "Usuario actualizado exitosamente", user)
}
