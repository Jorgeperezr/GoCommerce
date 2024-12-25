package utils

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de log: %v", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info registra mensajes de nivel informativo
func Info(message string) {
	infoLogger.Println(message)
}

// Error registra mensajes de nivel de error
func Error(message string) {
	errorLogger.Println(message)
}

// Fatal registra un error crítico y detiene la ejecución
func Fatal(message string) {
	errorLogger.Fatalln(message)
}
