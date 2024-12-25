package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User define el modelo para los usuarios en la base de datos
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);default:'user'"`
}

// BeforeCreate gorm hook para cifrar la contraseña antes de guardar
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword compara una contraseña plana con el hash almacenado
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// IsAdmin verifica si el usuario tiene rol de administrador
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
