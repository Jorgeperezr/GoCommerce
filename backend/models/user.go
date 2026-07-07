package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User define el modelo para los usuarios en la base de datos.
// La contraseña se serializa con json:"-" para no exponer nunca el hash.
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"type:varchar(20);default:'user'" json:"role"`
}

// SetPassword cifra y asigna la contraseña del usuario.
// Es el único punto donde se aplica bcrypt: quien crea o actualiza
// un usuario debe llamar a este método una sola vez.
func (u *User) SetPassword(plain string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
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
