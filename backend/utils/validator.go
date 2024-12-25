package utils

import (
	"errors"
	"regexp"
)

// ValidateEmail verifica si un correo electrónico tiene un formato válido
func ValidateEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("correo electrónico no válido")
	}
	return nil
}

// ValidatePassword verifica si una contraseña cumple con los requisitos mínimos
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	regex := `.*[A-Z].*`
	matched, err := regexp.MatchString(regex, password)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("la contraseña debe contener al menos una letra mayúscula")
	}
	return nil
}

// ValidateUsername verifica si un nombre de usuario es válido
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("el nombre de usuario debe tener al menos 3 caracteres")
	}
	return nil
}
