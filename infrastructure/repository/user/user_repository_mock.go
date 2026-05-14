package user

import (
	"context"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
)

type MockUserRepository struct{}

func (MockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user := &entity.User{
		ID:    id,
		Name:  "Lucas Chen",
		Email: "lucas.chen@example.com",
	}

	return user, nil
}

func (MockUserRepository) Save(context.Context, *entity.User) error {
	return nil
}
