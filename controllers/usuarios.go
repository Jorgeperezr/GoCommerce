package controllers

import (
	"encoding/json"
	"net/http"
)

// Usuario estructura del usuario
type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"`
}

// ObtenerUsuarios maneja la obtención de usuarios
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []Usuario
	// Simulación de datos
	usuarios = append(usuarios, Usuario{ID: 1, Nombre: "Admin", Correo: "admin@example.com", Rol: "Administrador"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// EliminarUsuario maneja la eliminación de un usuario
func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	// Lógica para eliminar un usuario
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado correctamente"))
}
