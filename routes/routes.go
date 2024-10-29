package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes registra todas as rotas http.
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	RegisterUserRoutes(api)
}
