package database

import (
	"log"
)

// ObtenerHistorial obtiene el historial de transacciones
func ObtenerHistorial() ([]Historial, error) {
	rows, err := DB.Query("SELECT id, producto, usuario, tipo, cantidad, fecha FROM historial_transacciones")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historial []Historial
	for rows.Next() {
		var h Historial
		if err := rows.Scan(&h.ID, &h.Producto, &h.Usuario, &h.Tipo, &h.Cantidad, &h.Fecha); err != nil {
			return nil, err
		}
		historial = append(historial, h)
	}
	return historial, nil
}
