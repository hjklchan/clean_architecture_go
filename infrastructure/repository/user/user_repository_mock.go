package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/user"
)

type MockUserRepository struct{}

func (MockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	user := &user.User{
		ID:    id,
		Name:  "Lucas Chen",
		Email: "lucas.chen@example.com",
	}

	return user, nil
}

func (MockUserRepository) Save(ctx context.Context, user *user.User) error {
	fmt.Println("====== infrastructure.repository.user.mock ======")
	fmt.Printf("get user data: %#v\n", user)
	fmt.Printf("save user successfully\n")

	return nil
}
