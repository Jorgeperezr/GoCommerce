package utils

import (
	"strconv"
	"time"
)

// ParseStringToInt convierte una cadena en un entero
func ParseStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// FormatDate formatea una fecha en formato YYYY-MM-DD
func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

// GetCurrentTimestamp obtiene la marca de tiempo actual
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// Contains verifica si un valor está presente en una lista de cadenas
func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
