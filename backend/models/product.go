package models

import (
	"fmt"
	"time"
)

// Product define el modelo para los productos en la base de datos
type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null;check:price > 0"`
	Stock       int       `gorm:"not null;default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// IsInStock verifica si el producto tiene stock disponible
func (p *Product) IsInStock() bool {
	return p.Stock > 0
}

// UpdateStock actualiza el stock del producto
func (p *Product) UpdateStock(quantity int) {
	if quantity >= 0 {
		p.Stock = quantity
	}
}

// DisplayPrice devuelve el precio formateado como string
func (p *Product) DisplayPrice() string {
	return fmt.Sprintf("$%.2f", p.Price)
}
