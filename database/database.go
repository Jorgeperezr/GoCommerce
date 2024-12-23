package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB es la conexión global de la base de datos
var DB *sql.DB

// ConectarDB inicializa la conexión a la base de datos
func ConectarDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "GoCommerce.db")
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos establecida correctamente")
	return nil
}

// CerrarDB cierra la conexión a la base de datos
func CerrarDB() error {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return fmt.Errorf("error al cerrar la base de datos: %v", err)
		}
		log.Println("Conexión a la base de datos cerrada correctamente")
	}
	return nil
}

// ObtenerProductos obtiene todos los productos de la base de datos
func ObtenerProductos() ([]Producto, error) {
	rows, err := DB.Query("SELECT id, nombre, precio, stock, categoria FROM productos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.Stock, &p.Categoria); err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

// EliminarProducto elimina un producto por ID
func EliminarProducto(id int) error {
	_, err := DB.Exec("DELETE FROM productos WHERE id = ?", id)
	return err
}
