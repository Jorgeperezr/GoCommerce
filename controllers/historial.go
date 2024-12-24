package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jorge/GoCommerce/database"
)

// HistorialHandler maneja la obtención del historial de transacciones
func HistorialHandler(w http.ResponseWriter, r *http.Request) {
	historial, err := database.ObtenerHistorial()
	if err != nil {
		log.Println("error al obtener historial desde el controlador:", err)
		http.Error(w, "error al obtener historial", http.StatusInternalServerError)
		return
	}

	if historial == nil {
		log.Println("historial is nil")
		http.Error(w, "no hay historial disponible", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(historial); err != nil {
		log.Println("error al codificar historial en JSON:", err)
		http.Error(w, "error interno", http.StatusInternalServerError)
	}
}
