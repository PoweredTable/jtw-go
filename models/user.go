package models

import "time"

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`          // Hashed password
	Role           string    `json:"role"`       // Adjust based on your user roles
	CreatedAt      time.Time `json:"created_at"` // Timestamp of when the user was created
	IsActive       bool      `json:"is_active"`  // Status of the user account
}
