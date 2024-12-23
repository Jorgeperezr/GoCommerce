package controllers

import (
	"GoCommerce/database"
	"encoding/json"
	"net/http"
)

// ObtenerHistorial maneja la solicitud de obtener el historial
func ObtenerHistorial(w http.ResponseWriter, r *http.Request) {
	historial, err := database.ObtenerHistorial()
	if err != nil {
		http.Error(w, "Error al obtener historial", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(historial)
}
