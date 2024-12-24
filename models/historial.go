package models

import "time"

// Historial representa un registro de actividad en el sistema
type Historial struct {
	ID      int    `json:"id"`
	Accion  string `json:"accion"`
	Usuario string `json:"usuario"`
	Fecha   time.Time `json:"fecha"`
}
