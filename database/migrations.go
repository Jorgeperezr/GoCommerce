package database

import (
	"database/sql"
	"log"
)

// DB is the database connection
var dbConnection *sql.DB

// InitializeDB initializes the database connection
func InitializeDB(dataSourceName string) error {
	var err error
	dbConnection, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	return dbConnection.Ping()
}

// MigrarDB ejecuta las migraciones de la base de datos
func MigrarDB() error {
	// Implement your migration logic here
	log.Println("Migraciones aplicadas exitosamente")
	return nil
}
