package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConectarDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./GoCommerce.db")
	if err != nil {
		log.Fatal("Error al conectar con SQLite:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	fmt.Println("Conexión a SQLite exitosa")
}
