package controllers

import (
	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"
	"net/http"
)

// Dashboard proporciona información general para administradores
func Dashboard(w http.ResponseWriter, r *http.Request) {
	var userCount int64
	var productCount int64

	if err := database.DB.Model(&models.User{}).Count(&userCount).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al obtener el total de usuarios", err.Error())
		return
	}

	if err := database.DB.Model(&models.Product{}).Count(&productCount).Error; err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Error al obtener el total de productos", err.Error())
		return
	}

	summary := map[string]interface{}{
		"total_users":    userCount,
		"total_products": productCount,
	}

	utils.SuccessResponse(w, "Datos del dashboard obtenidos correctamente", summary)
}
