package models

import (
	"GoCommerce/utils"
	"fmt"
)

// Producto representa un producto en el inventario.
type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID string
}

// Inventario representa una colección de productos.
type Inventario struct {
	Productos []Producto
}

// NuevoInventario crea un nuevo inventario vacío.
func NuevoInventario() *Inventario {
	return &Inventario{}
}

// RegistrarProducto agrega un producto al inventario.
func (inv *Inventario) RegistrarProducto(prod Producto) {
	inv.Productos = append(inv.Productos, prod)
	fmt.Println("Producto registrado exitosamente.")
}

// ListarProductos imprime todos los productos.
func (inv *Inventario) ListarProductos() {
	for _, prod := range inv.Productos {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d\n", prod.ID, prod.Nombre, prod.Precio, prod.Stock)
	}
}
