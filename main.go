/*
@autor: Jorge Pérez Rodríguez
@fecha: 30/10/2024
@descripcion: establece la base del backend para un sistema de e-commerce, definiendo las estructuras Producto y Pedido con métodos básicos.
Producto incluye un metodo para actualizar su stock, mientras que Pedido permite calcular el total de los productos y cambiar su estado entre opciones válidas.
La función main simula el flujo inicial al crear un producto y un pedido, ajustando el stock, calculando el total y actualizando el estado del pedido, proporcionando una estructura inicial para expandir el sistema de gestión de e-commerce.
*/
package main

import (
	"fmt"
	"time"
)

// Estructura Producto: representa los detalles de un producto en el inventario
type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID int
}

// Metodo para actualizar el stock de un producto
func (prod *Producto) ActualizarStock(cantidad int) {
	// Verificar que la cantidad no sea negativa
	if cantidad < 0 {
		fmt.Println("Error: la cantidad no puede ser negativa.")
		return
	}
	prod.Stock += cantidad
	fmt.Printf("Stock actualizado: %d unidades.\n", prod.Stock)
}

// Estructura Pedido: almacena los detalles de un pedido realizado
type Pedido struct {
	ID        int
	Fecha     time.Time
	Total     float64
	Estado    string
	Productos []Producto
}

// Metodo para calcular el total de un pedido en base a los precios de los productos
func (p *Pedido) CalcularTotal() {
	var total float64
	// Sumar los precios de todos los productos en el pedido
	for _, producto := range p.Productos {
		total += producto.Precio
	}
	p.Total = total
	fmt.Printf("Total del pedido calculado: %.2f\n", p.Total)
}

// Metodo para actualizar el estado de un pedido
func (p *Pedido) ActualizarEstado(nuevoEstado string) {
	// Definir estados validos
	estadosValidos := []string{"pendiente", "procesado", "enviado", "entregado"}
	esValido := false

	// Verificar que el nuevo estado sea valido
	for _, estado := range estadosValidos {
		if nuevoEstado == estado {
			esValido = true
			break
		}
	}
	if !esValido {
		fmt.Println("Error: estado no valido.")
		return
	}

	// Actualizar el estado si es valido
	p.Estado = nuevoEstado
	fmt.Printf("Estado del pedido actualizado a %s.\n", p.Estado)
}

// Funcion principal para inicializar el flujo del sistema
func main() {
	// Crear un producto de ejemplo
	producto := Producto{ID: 101, Nombre: "Laptop", Precio: 1200.00, Stock: 5, CategoriaID: 1}

	// Actualizar stock del producto
	producto.ActualizarStock(10)

	// Crear un pedido de ejemplo con una lista de productos
	pedido := Pedido{
		ID:        1001,
		Fecha:     time.Now(),
		Estado:    "pendiente",
		Productos: []Producto{producto},
	}

	// Calcular el total del pedido
	pedido.CalcularTotal()

	// Actualizar el estado del pedido
	pedido.ActualizarEstado("procesado")
}
