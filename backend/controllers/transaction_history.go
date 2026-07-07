package controllers

import (
	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTransactionHistory obtiene el historial de transacciones de un usuario específico
func GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var transactions []models.Transaction
	if err := database.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al obtener el historial de transacciones", err.Error())
		return
	}

	utils.SuccessResponse(w, "Historial de transacciones obtenido exitosamente", transactions)
}

// GetAllTransactions obtiene todas las transacciones
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	if err := database.DB.Find(&transactions).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al obtener las transacciones", err.Error())
		return
	}

	utils.SuccessResponse(w, "Todas las transacciones obtenidas exitosamente", transactions)
}
