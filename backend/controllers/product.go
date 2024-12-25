package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"

	"github.com/gorilla/mux"
)

// GetProducts obtiene todos los productos
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if result := database.DB.Find(&products); result.Error != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID obtiene un producto por su ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateProduct crea un nuevo producto
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price <= 0 {
		http.Error(w, "Nombre y precio son obligatorios y el precio debe ser positivo", http.StatusBadRequest)
		return
	}

	if result := database.DB.Create(&product); result.Error != nil {
		http.Error(w, "Error al crear producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct actualiza un producto por su ID
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if result := database.DB.Save(&product); result.Error != nil {
		http.Error(w, "Error al actualizar producto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct elimina un producto por su ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	if result := database.DB.Delete(&product); result.Error != nil {
		http.Error(w, "Error al eliminar producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado correctamente"})
}
