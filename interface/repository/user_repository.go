package repository

import (
	"context"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/user"
)

type UserRepository interface {
	GetByID(context.Context, uuid.UUID) (*user.User, error)
	Save(context.Context, *user.User) error
	// more
}
