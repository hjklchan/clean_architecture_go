package user

import (
	"time"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/user/value_object"
)

// Simple user entity
type User struct {
	ID               uuid.UUID
	Name             string
	Email            string
	Status           int
	DateOfBirth      value_object.DateOfBirth
	Password         value_object.Password
	PasswordExpiryAt value_object.PasswordExpiryAt
}

func NewUser(
	name,
	email string,
	dateOfBirth value_object.DateOfBirth,
	password value_object.Password,
	expiryDuration time.Duration,
) *User {
	passwordExpiryAt := time.Now().Add(expiryDuration)

	return &User{
		ID:               uuid.New(),
		Name:             name,
		Email:            email,
		DateOfBirth:      dateOfBirth,
		Password:         password,
		PasswordExpiryAt: value_object.NewPasswordExpiryAtFromTime(passwordExpiryAt),
	}
}

func (u *User) ChangeName(value string) {
	u.Name = value
}

func (u *User) ChangeEmail(value string) {
	u.Email = value
}

func (u *User) ResetPassword(value value_object.Password) {
	u.Password = value
}

func (u *User) CanLogin() bool {
	return !u.PasswordExpiryAt.IsExpired()
}
