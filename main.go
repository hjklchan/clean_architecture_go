package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
	"practices.com/clean_arch_go/usecase/user/get_user_by_id"
)

func main() {
	// Example for GetUserInfoById
	interactor := get_user_by_id.NewGetUserByIdInteractor(
		mockUserRepository{},
		GetUserByIdApiPresenter{},
	)

	interactor.Invoke(context.Background(), get_user_by_id.Input{ID: uuid.New()})
}

type mockUserRepository struct{}

func (mockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user := &entity.User{
		ID:    id,
		Name:  "Lucas Chen",
		Email: "lucas.chen@example.com",
	}

	return user, nil
}

func (mockUserRepository) Save(context.Context, *entity.User) error {
	return nil
}

// Implements output port
type GetUserByIdApiPresenter struct{}

func (p GetUserByIdApiPresenter) Present(output get_user_by_id.Output) error {
	return nil
}

func (p GetUserByIdApiPresenter) HandleError(error) error {
	fmt.Println("business error")

	return nil
}
