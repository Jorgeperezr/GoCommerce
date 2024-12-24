// Package models contains the data models for the GoCommerce application.
package models

// Usuario representa un usuario en el sistema
type Usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}
