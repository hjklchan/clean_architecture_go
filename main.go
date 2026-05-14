package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
	"practices.com/clean_arch_go/usecase/user/get_user_by_id"
)

func main() {
	// Repository Instance
	repository := mockUserRepository{}
	// Presenter Instance
	presenter := GetUserByIdApiPresenter{}

	// Example for GetUserInfoById
	interactor := get_user_by_id.NewGetUserByIdInteractor(
		repository,
		presenter,
	)

	err := interactor.Invoke(context.Background(), get_user_by_id.Input{ID: uuid.New()})
	if err != nil {
		fmt.Println("system error:", err)
		return
	}
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

type Response struct {
	ID    string
	Name  string
	Email string
}

// Implements output port
type GetUserByIdApiPresenter struct {
	ViewModel interface{}
}

func (p GetUserByIdApiPresenter) Present(output get_user_by_id.Output) error {
	fmt.Println("GetUserByIdApiPresenter.Present")
	fmt.Printf("Get the result which find the user by id: %#v\n", output)

	return nil
}

func (p GetUserByIdApiPresenter) HandleError(err error) error {
	fmt.Println("GetUserByIdApiPresenter.HandleError")
	fmt.Println("business error:", err)

	return nil
}
