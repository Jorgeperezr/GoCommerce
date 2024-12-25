package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse envía una respuesta JSON estándar
func JSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error al codificar la respuesta JSON: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}

// SuccessResponse envía una respuesta de éxito
func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	JSONResponse(w, http.StatusOK, message, data)
}

// ErrorResponse envía una respuesta de error
func ErrorResponse(w http.ResponseWriter, status int, message string, details interface{}) {
	JSONResponse(w, status, message, map[string]interface{}{
		"error":   true,
		"details": details,
	})
}
