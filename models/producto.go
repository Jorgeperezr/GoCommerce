package models

import "fmt"

type Producto struct {
	ID          int
	nombre      string
	precio      float64
	stock       int
	categoriaID string
}

func NuevoProducto(id int, nombre string, precio float64, stock int, categoriaID string) Producto {
	return Producto{
		ID:          id,
		nombre:      nombre,
		precio:      precio,
		stock:       stock,
		categoriaID: categoriaID,
	}
}

func (p *Producto) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Producto) SetPrecio(precio float64) {
	p.precio = precio
}

func (p *Producto) SetStock(stock int) {
	if stock < 0 {
		fmt.Println("Error: el stock no puede ser negativo.")
		return
	}
	p.stock = stock
}

func (p *Producto) SetCategoriaID(categoriaID string) {
	p.categoriaID = categoriaID
}

func (p *Producto) GetNombre() string {
	return p.nombre
}

func (p *Producto) MostrarDetalles() {
	fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d, Categoría: %s\n", p.ID, p.nombre, p.precio, p.stock, p.categoriaID)
}
