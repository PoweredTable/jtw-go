package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jtw-go/initializers"
	"time"
)

// GenerateJWT creates a new JWT token
func GenerateJWT(userID int, role string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    fmt.Sprint(userID),
		Subject:   role, // Store the role in the subject claim
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(initializers.JwtKey)
}

// ValidateJWT validates the JWT token and returns the claims
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
