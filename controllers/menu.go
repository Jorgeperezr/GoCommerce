package controllers

import (
	"GoCommerce/models"
	"GoCommerce/utils"
	"fmt"
)

func IniciarAplicacion() {
	fmt.Println("Bienvenido a GoCommerce")
	inventario := models.NuevoInventario()

	for {
		fmt.Println("\n1. Registrar Producto")
		fmt.Println("2. Buscar Producto")
		fmt.Println("3. Listar Productos")
		fmt.Println("4. Salir")
		opcion := utils.LeerEntrada("Seleccione una opción: ")

		switch opcion {
		case "1":
			utils.RegistrarProducto(inventario)
		case "2":
			utils.BuscarProducto(inventario)
		case "3":
			utils.ListarProductos(inventario)
		case "4":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
}
