package repository

import (
	"context"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
)

type UserRepository interface {
	GetByID(context.Context, uuid.UUID) (*entity.User, error)
	Save(context.Context, *entity.User) error
	// more
}
