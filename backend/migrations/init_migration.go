package migrations

import (
	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// InitMigration realiza la migración inicial de la base de datos y agrega usuarios predeterminados
func InitMigration() {
	// Migrar tablas
	database.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Transaction{})
	log.Println("Migraciones completadas correctamente")

	// Crear usuarios predeterminados
	users := []models.User{
		{Username: "admin", Password: hashPassword("1234"), Role: "admin"},
		{Username: "empleado1", Password: hashPassword("1234"), Role: "employee"},
	}

	for _, user := range users {
		var existingUser models.User
		if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
			database.DB.Create(&user)
			log.Printf("Usuario predeterminado creado: %s", user.Username)
		} else {
			log.Printf("El usuario %s ya existe", user.Username)
		}
	}
}

// hashPassword cifra una contraseña usando bcrypt
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error al cifrar la contraseña: %v", err)
	}
	return string(hashedPassword)
}
