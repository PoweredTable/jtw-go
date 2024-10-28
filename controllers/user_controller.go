package controllers

import (
	"github.com/gin-gonic/gin"
	"jtw-go/models"
	"jtw-go/usecases"
	"jtw-go/utils"
	"net/http"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
type userController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(useCase usecases.UserUseCase) UserController {
	return &userController{userUseCase: useCase}
}

func (uc *userController) Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration input"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	var user models.User

	user.Name = input.Name
	user.Phone = input.Phone
	user.Email = input.Email
	user.HashedPassword = hashedPassword

	err = uc.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hash": hashedPassword})
}

func (uc *userController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login input"})
		return
	}

	token, err := uc.userUseCase.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
