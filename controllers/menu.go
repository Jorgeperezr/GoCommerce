package controllers

import (
	"fmt"
	"net/http"
)

// Menu muestra las opciones principales del sistema
func Menu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido al sistema GoCommerce")
	fmt.Fprintln(w, "1. Gestión de Productos")
	fmt.Fprintln(w, "2. Historial de Transacciones")
	fmt.Fprintln(w, "3. Cerrar Sesión")
}
