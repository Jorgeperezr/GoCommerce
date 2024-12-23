package main

import (
	"GoCommerce/api"
	"GoCommerce/database"
	"log"
	"net/http"
)

func main() {
	// Conectar a la base de datos
	if err := database.ConectarDB(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer database.CerrarDB()

	// Configurar rutas
	api.ConfigurarRutas()

	// Iniciar servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
