package utils

import "log"

// ManejarError registra y muestra errores
func ManejarError(err error, mensaje string) {
	if err != nil {
		log.Printf("%s: %v", mensaje, err)
	}
	return
}
