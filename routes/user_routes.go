package routes

import (
	"github.com/gin-gonic/gin"
	"jtw-go/controllers"
	"jtw-go/initializers"
	"jtw-go/middleware"
	"jtw-go/repositories"
	"jtw-go/usecases"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRepo := repositories.NewUserRepository(initializers.DB)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	rg.POST("/login", userController.Login)

	rg.POST("/register",
		middleware.JWTAuthMiddleware,
		func(c *gin.Context) { middleware.RoleRequired("admin", c) },
		userController.Register,
	)
}
