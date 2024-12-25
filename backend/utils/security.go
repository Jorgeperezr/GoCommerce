package utils

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims estructura para almacenar los datos del token
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GetUserClaims obtiene las reclamaciones del token desde el contexto
func GetUserClaims(r *http.Request) (*Claims, bool) {
	claims, ok := r.Context().Value("user").(*Claims)
	return claims, ok
}

// GenerateToken genera un nuevo token JWT para un usuario autenticado
func GenerateToken(username, role string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Println("SECRET_KEY no definido en las variables de entorno")
		return "", jwt.ErrSignatureInvalid
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error al firmar el token: %v", err)
		return "", err
	}

	return tokenString, nil
}
