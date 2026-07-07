package controllers

import (
	"encoding/json"
	"net/http"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateOrderRequest datos que el cliente puede enviar al crear un pedido.
// El precio total y el usuario NUNCA vienen del cliente: se calculan en el servidor.
type CreateOrderRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

// CreateOrder crea un nuevo pedido validando stock y calculando el total
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetUserClaims(r)
	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "No autorizado", nil)
		return
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}
	if req.ProductID == 0 || req.Quantity <= 0 {
		utils.ErrorResponse(w, http.StatusBadRequest, "product_id y una cantidad mayor que cero son obligatorios", nil)
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Usuario no encontrado", nil)
		return
	}

	var order models.Order
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var product models.Product
		if err := tx.First(&product, req.ProductID).Error; err != nil {
			return err
		}
		if product.Stock < req.Quantity {
			return gorm.ErrInvalidData
		}

		product.Stock -= req.Quantity
		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		order = models.Order{
			UserID:     user.ID,
			ProductID:  product.ID,
			Quantity:   req.Quantity,
			TotalPrice: product.Price * float64(req.Quantity),
			Status:     utils.OrderStatusPending,
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			UserID: user.ID,
			Amount: order.TotalPrice,
			Status: utils.OrderStatusCompleted,
		}
		return tx.Create(&transaction).Error
	})

	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "No se pudo crear el pedido (¿producto inexistente o stock insuficiente?)", err.Error())
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

// UpdateOrderRequest solo permite cambiar el estado del pedido
type UpdateOrderRequest struct {
	Status string `json:"status"`
}

// UpdateOrder actualiza el estado de un pedido por su ID
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Pedido no encontrado", err.Error())
		return
	}

	var req UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	validStatus := []string{utils.OrderStatusPending, utils.OrderStatusCompleted, utils.OrderStatusCancelled}
	if !utils.Contains(validStatus, req.Status) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Estado de pedido inválido", nil)
		return
	}

	order.Status = req.Status
	if err := database.DB.Save(&order).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al actualizar el pedido", err.Error())
		return
	}
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
