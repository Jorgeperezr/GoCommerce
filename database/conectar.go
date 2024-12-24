package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConectarDB() error {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ecommerce")
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("Conexión a la base de datos exitosa")
	return nil
}
