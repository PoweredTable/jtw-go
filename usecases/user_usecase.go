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

func (u *userUseCase) Register(user models.User) error {
	return u.userRepo.CreateUser(user)
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.HashedPassword) {
		return "", fmt.Errorf("invalid credentials")
	}

	return utils.GenerateJWT(user.ID, user.Role)
}
