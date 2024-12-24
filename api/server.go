package api

import (
	"log"
	"net/http"
)

// ConfigurarRutasServidor configura las rutas del servidor HTTP
func ConfigurarRutasServidor() {
	// Aquí se configuran las rutas
}

// IniciarServidor inicializa el servidor HTTP
func IniciarServidor() {
	ConfigurarRutasServidor()
	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
