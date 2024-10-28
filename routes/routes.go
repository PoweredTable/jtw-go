package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	RegisterUserRoutes(api)
}
