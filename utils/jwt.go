package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jtw-go/initializers"
	"time"
)

// GenerateJWT cria um novo token JWT utilizando o ID do usuário
// e sua Role e retorna o token gerado ou um erro.
func GenerateJWT(userID int, role string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    fmt.Sprint(userID),
		Subject:   role, // Armazena a role na claim subject
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(initializers.JwtKey)
}

// ValidateJWT valida o token JWT e retorna as claims.
func ValidateJWT(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return initializers.JwtKey, nil
	})
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
