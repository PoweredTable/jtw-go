package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"jtw-go/utils"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("userID", claims.Issuer)
	c.Set("role", claims.Subject)
	c.Next()
}

// RoleRequired verifica se o usuário possui a role necessária.
func RoleRequired(role string, c *gin.Context) {
	// Adquire a role direto do contexto e verifica se existe
	userRole, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	// Compara a role do usuário logado com a role exigida
	if userRole != role {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
		return
	}
	c.Next()
}
