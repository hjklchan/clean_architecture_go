package create_user

import (
	"context"
	"time"

	"practices.com/clean_arch_go/domain/user"
	"practices.com/clean_arch_go/domain/user/value_object"
)

type CreateUserUseCase struct {
	repo           user.UserRepository
	passwordHasher value_object.PasswordHasher
}

func NewCreateUserUseCase(
	repo user.UserRepository,
	passwordHasher value_object.PasswordHasher,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo,
		passwordHasher,
	}
}

func (uc *CreateUserUseCase) Invoke(ctx context.Context, in Input) error {
	p, err := value_object.NewPasswordFromPlainText(
		in.Password,
		uc.passwordHasher,
	)
	if err != nil {
		return err
	}

	const expiryDur = time.Hour * 24
	ent := user.NewUser(in.Name, in.Email, *p, expiryDur)

	return uc.repo.Save(ctx, ent)
}
