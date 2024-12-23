// Package utils proporciona funciones de validación.
package utils

import (
	"errors"
	"strings"
)

// ValidarEntrada valida que una entrada no esté vacía.
func ValidarEntrada(entrada, tipo string) (bool, error) {
	entrada = strings.TrimSpace(entrada)
	if entrada == "" {
		return false, errors.New("la entrada no puede estar vacía")
	}
	return true, nil
}
