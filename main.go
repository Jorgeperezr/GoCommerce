/*
@autor: Jorge Pérez Rodríguez
@fecha: 20/11/2024
@descripcion: Se añadió funcionalidad de registro, consulta, actualización y eliminación de productos mediante un menú interactivo, con validación de entradas y mensajes de confirmación para acciones realizadas..
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// estructura producto: representa un producto del inventario
type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID string
}

// metodo para actualizar el stock de un producto
func (prod *Producto) ActualizarStock(cantidad int) {
	if cantidad < 0 {
		fmt.Println("error: la cantidad no puede ser negativa.")
		return
	}
	prod.Stock += cantidad
	fmt.Printf("stock actualizado a %d unidades.\n", prod.Stock)
}

// estructura para gestionar el inventario de productos
type Inventario struct {
	Productos []Producto
}

// metodo para agregar un nuevo producto al inventario
func (inv *Inventario) AgregarProducto(producto Producto) {
	inv.Productos = append(inv.Productos, producto)
	fmt.Println("producto registrado con exito:")
	fmt.Printf("id: %d, nombre: %s, precio: %.2f, stock: %d, categoria: %s\n", producto.ID, producto.Nombre, producto.Precio, producto.Stock, producto.CategoriaID)
}

// metodo para buscar un producto por id
func (inv *Inventario) BuscarProductoPorID(id int) *Producto {
	for i, producto := range inv.Productos {
		if producto.ID == id {
			return &inv.Productos[i]
		}
	}
	return nil
}

// metodo para eliminar un producto por id
func (inv *Inventario) EliminarProducto(id int) {
	for i, producto := range inv.Productos {
		if producto.ID == id {
			inv.Productos = append(inv.Productos[:i], inv.Productos[i+1:]...)
			fmt.Println("producto eliminado con exito.")
			return
		}
	}
	fmt.Println("error: producto no encontrado.")
}

// metodo para mostrar el menu de opciones
func MostrarMenu() {
	fmt.Println("\n.................  Menu  .................")
	fmt.Println("1. Registrar un nuevo producto")
	fmt.Println("2. Buscar un producto por id")
	fmt.Println("3. Actualizar informacion de un producto")
	fmt.Println("4. Actualizar stock de un producto")
	fmt.Println("5. Eliminar un producto")
	fmt.Println("6. Salir")
}

// funcion para leer entradas del usuario
func LeerEntrada(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	entrada, _ := reader.ReadString('\n')
	return strings.TrimSpace(entrada)
}

// funcion principal
func main() {
	inventario := Inventario{}
	idActual := 1 // contador para asignar ids unicos

	for {
		MostrarMenu()
		opcionStr := LeerEntrada("seleccione una opcion: ")
		opcion, err := strconv.Atoi(opcionStr)
		if err != nil {
			fmt.Println("error: opcion invalida.")
			continue
		}

		switch opcion {
		case 1: // registrar un nuevo producto
			nombre := LeerEntrada("ingrese el nombre del producto: ")
			precioStr := LeerEntrada("ingrese el precio del producto: ")
			precio, err := strconv.ParseFloat(precioStr, 64)
			if err != nil {
				fmt.Println("error: el precio debe ser un numero.")
				continue
			}

			stockStr := LeerEntrada("ingrese el stock del producto: ")
			stock, err := strconv.Atoi(stockStr)
			if err != nil {
				fmt.Println("error: el stock debe ser un numero entero.")
				continue
			}

			categoria := LeerEntrada("ingrese la categoria del producto: ")

			producto := Producto{
				ID:          idActual,
				Nombre:      nombre,
				Precio:      precio,
				Stock:       stock,
				CategoriaID: categoria,
			}
			inventario.AgregarProducto(producto)
			idActual++

		case 2: // buscar un producto por id
			idStr := LeerEntrada("ingrese el id del producto: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("error: el id debe ser un numero entero.")
				continue
			}

			producto := inventario.BuscarProductoPorID(id)
			if producto != nil {
				fmt.Printf("producto encontrado: id: %d, nombre: %s, precio: %.2f, stock: %d, categoria: %s\n", producto.ID, producto.Nombre, producto.Precio, producto.Stock, producto.CategoriaID)
			} else {
				fmt.Println("error: producto no encontrado.")
			}

		case 3: // actualizar informacion de un producto
			idStr := LeerEntrada("ingrese el id del producto a actualizar: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("error: el id debe ser un numero entero.")
				continue
			}

			producto := inventario.BuscarProductoPorID(id)
			if producto != nil {
				producto.Nombre = LeerEntrada("ingrese el nuevo nombre del producto: ")
				precioStr := LeerEntrada("ingrese el nuevo precio del producto: ")
				precio, err := strconv.ParseFloat(precioStr, 64)
				if err != nil {
					fmt.Println("error: el precio debe ser un numero.")
					continue
				}
				producto.Precio = precio

				stockStr := LeerEntrada("ingrese el nuevo stock del producto: ")
				stock, err := strconv.Atoi(stockStr)
				if err != nil {
					fmt.Println("error: el stock debe ser un numero entero.")
					continue
				}
				producto.Stock = stock

				producto.CategoriaID = LeerEntrada("ingrese la nueva categoria del producto: ")
				fmt.Println("producto actualizado con exito:")
				fmt.Printf("id: %d, nombre: %s, precio: %.2f, stock: %d, categoria: %s\n", producto.ID, producto.Nombre, producto.Precio, producto.Stock, producto.CategoriaID)
			} else {
				fmt.Println("error: producto no encontrado.")
			}

		case 4: // actualizar stock de un producto
			idStr := LeerEntrada("ingrese el id del producto para actualizar el stock: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("error: el id debe ser un numero entero.")
				continue
			}

			producto := inventario.BuscarProductoPorID(id)
			if producto != nil {
				cantidadStr := LeerEntrada("ingrese la cantidad para actualizar el stock: ")
				cantidad, err := strconv.Atoi(cantidadStr)
				if err != nil {
					fmt.Println("error: la cantidad debe ser un numero entero.")
					continue
				}
				producto.ActualizarStock(cantidad)
			} else {
				fmt.Println("error: producto no encontrado.")
			}

		case 5: // eliminar un producto
			idStr := LeerEntrada("ingrese el id del producto a eliminar: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("error: el id debe ser un numero entero.")
				continue
			}
			inventario.EliminarProducto(id)

		case 6: // salir
			fmt.Println("saliendo del sistema...")
			return

		default:
			fmt.Println("error: opcion no valida.")
		}
	}
}
