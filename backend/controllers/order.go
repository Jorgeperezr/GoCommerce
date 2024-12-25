package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"github.com/gorilla/mux"
)

// CreateOrder crea un nuevo pedido
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	if err := database.DB.Create(&order).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al crear el pedido", err.Error())
		return
	}

	utils.SuccessResponse(w, "Pedido creado exitosamente", order)
}

// GetOrder obtiene un pedido por su ID
func GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Pedido no encontrado", err.Error())
		return
	}

	utils.SuccessResponse(w, "Pedido obtenido exitosamente", order)
}

// GetAllOrders obtiene todos los pedidos
func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al obtener los pedidos", err.Error())
		return
	}

	utils.SuccessResponse(w, "Pedidos obtenidos exitosamente", orders)
}

// UpdateOrder actualiza un pedido por su ID
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Pedido no encontrado", err.Error())
		return
	}

	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	database.DB.Model(&order).Updates(updatedOrder)
	utils.SuccessResponse(w, "Pedido actualizado exitosamente", order)
}

// DeleteOrder elimina un pedido por su ID
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	if err := database.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al eliminar el pedido", err.Error())
		return
	}

	utils.SuccessResponse(w, "Pedido eliminado exitosamente", nil)
}
