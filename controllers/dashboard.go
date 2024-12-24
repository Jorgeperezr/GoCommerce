package controllers

import (
	"net/http"
)

// HandleDashboard maneja la página principal después del login
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "acceso no autorizado", http.StatusUnauthorized)
		return
	}
	username := cookie.Value

	if username == "admin" {
		w.Write([]byte("bienvenido administrador"))
	} else if username == "usuario" {
		w.Write([]byte("bienvenido empleado"))
	} else {
		http.Error(w, "acceso no autorizado", http.StatusUnauthorized)
	}
}
