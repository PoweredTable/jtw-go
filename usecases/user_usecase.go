package usecases

import (
	"fmt"
	"jtw-go/models"
	"jtw-go/repositories"
	"jtw-go/utils"
)

type UserUseCase interface {
	Login(email, password string) (string, error)
	Register(user models.User) error
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return &userUseCase{userRepo: repo}
}

// Register cria um novo usuário no banco de dados.
func (u *userUseCase) Register(user models.User) error {
	return u.userRepo.CreateUser(user)
}

// Login busca o usuário pelo email, verifica a senha hashed e retorna um token JWT.
func (u *userUseCase) Login(email, password string) (string, error) {
	// Busca o usuário pelo seu email
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	// Caso encontrado, verifica a senha com a senha hashed do usuário
	if !utils.CheckPasswordHash(password, user.HashedPassword) {
		return "", fmt.Errorf("invalid credentials")
	}
	// Gera um token JWT com ID e Role se atendido todas as condições
	return utils.GenerateJWT(user.ID, user.Role)
}
