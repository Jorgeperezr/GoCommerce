package controllers

import (
	"GoCommerce/database"
	"GoCommerce/models"
	"encoding/json"
	"net/http"
)

// ObtenerProductos devuelve todos los productos.
func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nombre, precio, stock, categoria FROM productos")
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var productos []models.Producto
	for rows.Next() {
		var p models.Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.Stock, &p.Categoria); err != nil {
			http.Error(w, "Error al escanear producto", http.StatusInternalServerError)
			return
		}
		productos = append(productos, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error al iterar sobre los productos", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(productos); err != nil {
		http.Error(w, "Error al codificar productos a JSON", http.StatusInternalServerError)
		return
	}
}

func ObtenerProductosHandler(w http.ResponseWriter, r *http.Request) {
	// Your product retrieval logic here
}
