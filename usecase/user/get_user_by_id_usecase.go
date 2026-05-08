package user

import (
	"context"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
	"practices.com/clean_arch_go/interface/repository"
)

type GetUserByIdUseCase struct {
	repo repository.UserRepository
}

func NewGetUserByIdUseCase(repo repository.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		repo,
	}
}

func (uc *GetUserByIdUseCase) Invoke(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	ent, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ent, nil
}
