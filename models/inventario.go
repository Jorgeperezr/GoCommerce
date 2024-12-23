// Package models define la estructura del inventario.
package models

// Inventario representa una colección de productos.
type Inventario struct {
	Productos []Producto
}

// NuevoInventario crea un nuevo inventario vacío.
func NuevoInventario() *Inventario {
	return &Inventario{
		Productos: []Producto{},
	}
}

// AgregarProducto agrega un nuevo producto al inventario.
func (i *Inventario) AgregarProducto(p Producto) {
	i.Productos = append(i.Productos, p)
}
