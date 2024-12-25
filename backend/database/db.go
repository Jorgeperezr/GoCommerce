package database

import (
	"fmt"
	"log"
	"os"

	"GoCommerce/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB establece la conexión con la base de datos
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	Dialect  string
}

func loadConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Dialect:  os.Getenv("DB_DIALECT"),
	}
}

func ConnectDB() error {
	config := loadConfig()
	var err error

	if config.Dialect == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host, config.User, config.Password, config.DBName, config.Port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		// Predeterminado a SQLite si no hay otro dialecto
		DB, err = gorm.Open(sqlite.Open("gocommerce.db"), &gorm.Config{})
	}

	if err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	log.Println("Base de datos conectada correctamente")

	if err := DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		return fmt.Errorf("error al migrar la base de datos: %v", err)
	}

	log.Println("Migración de base de datos completada")
	return nil
}
