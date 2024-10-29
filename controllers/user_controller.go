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

// Register recebe um input JSON através do gin.Context e tenta registrar o usuário.
func (uc *userController) Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// Valida o input de dados
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration input"})
		return
	}
	// Tenta criar um hash da senha do usuário
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

	// Tenta registrar o usuário
	err = uc.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login recebe um input JSON através do gin.Context e tenta realizar o login do usuário.
func (uc *userController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// Valida o input de dados
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login input"})
		return
	}
	// Tenta logar o usuário
	token, err := uc.userUseCase.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
