package controllers

import (
	"GoCommerce/models"
	"GoCommerce/utils"
	"fmt"
)

func Login(usuario string, contraseña string) bool {
	user := models.ObtenerUsuario(usuario)
	if user != nil && user.Contraseña == utils.HashPassword(contraseña) {
		fmt.Println("Inicio de sesión exitoso")
		return true
	}
	fmt.Println("Usuario o contraseña incorrecta")
	return false
}
