/*
@descripcion: proporciona funciones para hashing y validacion de entradas seguras.
*/

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

// funcion para generar un hash de una contraseña
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// funcion para validar contraseñas seguras
func ValidarContraseña(password string) bool {
	// contraseña debe tener al menos 8 caracteres, una mayuscula y un numero
	patron := `^(?=.*[A-Z])(?=.*\d)[A-Za-z\d@$!%*?&]{8,}$`
	return regexp.MustCompile(patron).MatchString(password)
}

// funcion para validar entradas seguras
func ValidarEntradaSegura(input string) bool {
	patron := `^[a-zA-Z0-9_@.-]+$`
	return regexp.MustCompile(patron).MatchString(input)
}
