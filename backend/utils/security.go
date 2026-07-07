package utils

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// contextKey es un tipo propio para las claves de contexto y evitar colisiones
type contextKey string

// UserContextKey clave bajo la cual el middleware guarda los claims del usuario
const UserContextKey contextKey = "user"

// Claims estructura para almacenar los datos del token
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GetUserClaims obtiene las reclamaciones del token desde el contexto
func GetUserClaims(r *http.Request) (*Claims, bool) {
	claims, ok := r.Context().Value(UserContextKey).(*Claims)
	return claims, ok
}

// WithUserClaims devuelve un contexto con los claims del usuario adjuntos
func WithUserClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, UserContextKey, claims)
}

// GenerateToken genera un nuevo token JWT para un usuario autenticado
func GenerateToken(username, role string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Println("SECRET_KEY no definido en las variables de entorno")
		return "", errors.New("SECRET_KEY no configurada")
	}

	expirationTime := time.Now().Add(TokenExpirationHours * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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

// ParseToken valida un token JWT y devuelve sus claims
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inesperado")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}
