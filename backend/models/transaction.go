package models

import (
	"time"
)

// Transaction representa una transacción en el sistema
type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName establece el nombre de la tabla para GORM
func (Transaction) TableName() string {
	return "transactions"
}
