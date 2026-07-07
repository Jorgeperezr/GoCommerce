package migrations

import (
	"log"

	"GoCommerce/backend/database"
	"GoCommerce/backend/models"
	"GoCommerce/backend/utils"
)

// InitMigration crea los usuarios predeterminados si no existen.
// Las tablas ya fueron migradas por database.ConnectDB.
func InitMigration() {
	defaults := []struct {
		Username string
		Password string
		Role     string
	}{
		{Username: "admin", Password: "Admin1234", Role: utils.RoleAdmin},
		{Username: "empleado1", Password: "Empleado1234", Role: utils.RoleUser},
	}

	for _, d := range defaults {
		var existingUser models.User
		if err := database.DB.Where("username = ?", d.Username).First(&existingUser).Error; err == nil {
			log.Printf("El usuario %s ya existe", d.Username)
			continue
		}

		user := models.User{Username: d.Username, Role: d.Role}
		if err := user.SetPassword(d.Password); err != nil {
			log.Printf("Error al cifrar la contraseña de %s: %v", d.Username, err)
			continue
		}
		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("Error al crear el usuario %s: %v", d.Username, err)
			continue
		}
		log.Printf("Usuario predeterminado creado: %s", d.Username)
	}
}
