package main

import (
	"log"
	"net/http"

	"github.com/jorge/GoCommerce/database"
	"github.com/jorge/GoCommerce/routes"
)

func main() {
	// Conectar a la base de datos
	if err := database.ConectarDB(); err != nil {
		log.Fatal("no se pudo conectar a la base de datos:", err)
	}
	defer database.CerrarDB()

	// Aplicar migraciones
	if err := database.MigrarDB(); err != nil {
		log.Fatal("no se pudieron aplicar las migraciones:", err)
	}

	// Configurar rutas
	routes.ConfigurarRutas()

	// Iniciar servidor
	log.Println("servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
