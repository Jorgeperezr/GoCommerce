package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"
)

// RegisterRequest datos esperados para crear un nuevo usuario
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register permite crear un nuevo usuario
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	if err := utils.ValidateUsername(req.Username); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Nombre de usuario inválido", err.Error())
		return
	}
	if err := utils.ValidatePassword(req.Password); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Contraseña inválida", err.Error())
		return
	}

	// El rol siempre es "user": los administradores se crean desde la
	// migración inicial o por otro administrador, nunca por auto-registro.
	user := models.User{Username: req.Username, Role: utils.RoleUser}
	if err := user.SetPassword(req.Password); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al cifrar la contraseña", err.Error())
		return
	}

	if result := database.DB.Create(&user); result.Error != nil {
		utils.ErrorResponse(w, http.StatusConflict, "Error al crear el usuario", result.Error.Error())
		return
	}

	utils.SuccessResponse(w, "Usuario creado exitosamente", user)
}
