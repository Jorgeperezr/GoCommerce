package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// Register permite crear un nuevo usuario
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	if user.Username == "" || user.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Nombre de usuario y contraseña son obligatorios", nil)
		return
	}

	// Cifrar la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al cifrar la contraseña", err.Error())
		return
	}
	user.Password = string(hashedPassword)

	// Crear usuario en la base de datos
	if result := database.DB.Create(&user); result.Error != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al crear el usuario", result.Error.Error())
		return
	}

	utils.SuccessResponse(w, "Usuario creado exitosamente", user)
}
