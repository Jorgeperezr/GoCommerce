package controllers

import (
	"net/http"
)

// HandleDashboard maneja la página principal después del login
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_token")
	username := cookie.Value

	if username == "admin" {
		w.Write([]byte("bienvenido administrador"))
	} else if username == "usuario" {
		w.Write([]byte("bienvenido empleado"))
	} else {
		http.Error(w, "acceso no autorizado", http.StatusUnauthorized)
	}
}
