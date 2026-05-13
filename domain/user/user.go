package user

import "github.com/google/uuid"

// Simple user entity
type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}
}
