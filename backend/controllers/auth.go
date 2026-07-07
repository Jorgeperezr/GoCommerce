package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "El nombre de usuario y la contraseña son obligatorios", http.StatusBadRequest)
		return
	}

	var user models.User
	if result := database.DB.Where("username = ?", creds.Username).First(&user); result.Error != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	if !user.CheckPassword(creds.Password) {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
