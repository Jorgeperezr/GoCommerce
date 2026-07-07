package models

import (
	"time"
)

// Order representa un pedido en el sistema
type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName establece el nombre de la tabla para GORM
func (Order) TableName() string {
	return "orders"
}
