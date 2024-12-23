/*
@descripcion: manejo centralizado de errores en el sistema.
*/

package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// funcion para manejar errores genericos
func ManejarError(err error, mensaje string) error {
	if err != nil {
		log.Printf("error: %s - %v", mensaje, err)
		return errors.New(mensaje)
	}
	return nil
}

// funcion para responder con error http
func ResponderConError(w http.ResponseWriter, status int, mensaje string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("error: %s", mensaje)))
}

// funcion para validar errores fatales
func FatalError(err error, mensaje string) {
	if err != nil {
		log.Fatalf("error fatal: %s - %v", mensaje, err)
	}
}
