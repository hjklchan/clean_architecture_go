package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetByID(context.Context, uuid.UUID) (*User, error)
	Save(context.Context, *User) error
	// more
}
