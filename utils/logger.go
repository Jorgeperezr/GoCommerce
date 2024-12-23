/*
@descripcion: proporciona funciones para el registro de eventos y errores.
*/

package utils

import (
	"log"
	"os"
)

// funcion para inicializar un archivo de log
func IniciarLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error al crear archivo de log: %v", err)
	}
	log.SetOutput(file)
}

// funcion para registrar un mensaje informativo
func Info(mensaje string) {
	log.Printf("info: %s", mensaje)
}

// funcion para registrar una advertencia
func Advertencia(mensaje string) {
	log.Printf("advertencia: %s", mensaje)
}

// funcion para registrar un error
func Error(mensaje string) {
	log.Printf("error: %s", mensaje)
}
