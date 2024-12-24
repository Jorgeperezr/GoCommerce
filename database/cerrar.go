package database

func CerrarDB() {
	if DB != nil {
		DB.Close()
	}
}
