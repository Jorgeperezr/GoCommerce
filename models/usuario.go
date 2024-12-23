package models

// Usuario representa la estructura de un usuario en la base de datos
type Usuario struct {
	ID     int
	Nombre string
	Correo string
	Rol    string
}
