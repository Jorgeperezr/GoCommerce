package controllers

import (
	"GoCommerce/database"
	"encoding/json"
	"net/http"
	"strconv"
)

// ObtenerProductos maneja la solicitud de obtener productos
func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	productos, err := database.ObtenerProductos()
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(productos)
}

// EliminarProducto maneja la solicitud de eliminar un producto
func EliminarProducto(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := database.EliminarProducto(id); err != nil {
		http.Error(w, "Error al eliminar producto", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto eliminado correctamente"))
}
