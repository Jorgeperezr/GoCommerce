package config

import (
	"log"
	"os"
)

// Config estructura que almacena las variables de configuración clave
type Config struct {
	AppPort   string
	SecretKey string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBDialect string
}

// LoadConfig carga las variables de entorno necesarias
func LoadConfig() *Config {
	config := &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		SecretKey: getEnv("SECRET_KEY", "default_secret"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASSWORD", ""),
		DBName:    getEnv("DB_NAME", "gocommerce"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBDialect: getEnv("DB_DIALECT", "sqlite"),
	}

	validateCriticalEnv(config)

	log.Println("Configuración cargada correctamente")
	return config
}

// getEnv obtiene una variable de entorno con un valor por defecto
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// validateCriticalEnv verifica que las variables críticas estén configuradas
func validateCriticalEnv(config *Config) {
	if config.SecretKey == "default_secret" {
		log.Fatal("ERROR: SECRET_KEY no está configurada correctamente")
	}
	if config.DBPass == "" {
		log.Fatal("ERROR: DB_PASSWORD no está configurada")
	}
}
