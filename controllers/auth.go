package controllers

import (
	"net/http"
)

// handlelogin maneja el inicio de sesión de usuarios
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("usuario")
	password := r.FormValue("contrasena")

	if username == "admin" && password == "1234" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	} else if username == "usuario" && password == "1234" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
}

// handlelogout maneja el cierre de sesión de usuarios
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
