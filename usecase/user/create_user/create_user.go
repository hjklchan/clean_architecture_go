package create_user

import (
	"context"

	"practices.com/clean_arch_go/entity"
	"practices.com/clean_arch_go/interface/repository"
)

type CreateUserUseCase struct {
	repo repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo,
	}
}

func (uc *CreateUserUseCase) Invoke(ctx context.Context, name, email string) error {
	ent := entity.NewUser(name, email)

	return uc.repo.Save(ctx, ent)
}
