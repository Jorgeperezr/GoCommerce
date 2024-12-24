package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB representa la conexión global a la base de datos

// ConectarDB inicializa la conexión a la base de datos
func ConectarDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./GoCommerce.db")
	if err != nil {
		log.Println("error al conectar con la base de datos:", err)
		return err
	}

	log.Println("conexión exitosa a la base de datos")
	return nil
}

// CerrarDB cierra la conexión a la base de datos
func CerrarDB() {
	if DB != nil {
		DB.Close()
		log.Println("conexión a la base de datos cerrada")
	}
}
