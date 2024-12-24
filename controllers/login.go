package controllers

import (
	"GoCommerce/database"
	"GoCommerce/models"
	"encoding/json"
	"net/http"
)

// AuthHandler maneja la autenticación de usuarios
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Usuario
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "error al leer los datos de autenticación", http.StatusBadRequest)
		return
	}

	usuario, err := database.ObtenerUsuarioPorCorreo(creds.Email)
	if err != nil {
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}

	if usuario.ID == 0 || usuario.Password != creds.Password {
		http.Error(w, "credenciales inválidas", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "autenticación exitosa",
		"user":    usuario.Nombre,
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Your login logic here
}
