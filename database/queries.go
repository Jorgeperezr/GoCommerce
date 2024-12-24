package database

import (
	"GoCommerce/models"
	"database/sql"
	"log"

	"github.com/jorge/GoCommerce/models"
)

// DB is a global variable to hold the database connection
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("error al hacer ping a la base de datos:", err)
	}
}

// ObtenerUsuarioPorCorreo busca un usuario por su correo electrónico
func ObtenerUsuarioPorCorreo(email string) (models.Usuario, error) {
	var usuario models.Usuario
	query := "SELECT id, nombre, email, password FROM usuarios WHERE email = ?"
	err := DB.QueryRow(query, email).Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return usuario, nil
		}
		log.Println("Error al obtener usuario por correo:", err)
		return usuario, err
	}
	return usuario, nil
}

// ObtenerHistorial recupera todos los registros del historial
func ObtenerHistorial() ([]models.Historial, error) {
	query := "SELECT id, accion, usuario, fecha FROM historial"
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("error al obtener historial:", err)
		return nil, err
	}
	defer rows.Close()

	var historial []models.Historial
	for rows.Next() {
		var registro models.Historial
		if err := rows.Scan(&registro.ID, &registro.Accion, &registro.Usuario, &registro.Fecha); err != nil {
			log.Println("error al leer registro de historial:", err)
			return nil, err
		}
		historial = append(historial, registro)
	}

	return historial, nil
}
