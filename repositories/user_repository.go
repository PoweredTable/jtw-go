package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"jtw-go/models"
	"log"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// CreateUser cria um novo usuário no banco de dados.
func (r *userRepository) CreateUser(user models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO \"user\" (name, phone, email, password_hash) VALUES ($1, $2, $3, $4)",
		user.Name, user.Phone, user.Email, user.HashedPassword,
	)
	return err
}

// GetUserByEmail retorna um usuário pelo seu email.
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	stmt, err := r.db.Prepare("SELECT id, email, password_hash, role FROM \"user\" WHERE email = $1")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.HashedPassword, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Retorna um erro caso o usuário não exista
			return nil, fmt.Errorf("user does not exists")
		}
		log.Println("Error fetching user:", err)
		return nil, err
	}
	return &user, nil
}
