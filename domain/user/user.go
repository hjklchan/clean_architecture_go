package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/user/value_object"
	"practices.com/clean_arch_go/shared/shared_value_object"
)

const PasswordExpiryDuration = time.Hour * 24 * 30

// Simple user entity
type User struct {
	id               shared_value_object.Id
	name             string
	email            string
	status           int
	dateOfBirth      value_object.DateOfBirth
	password         value_object.Password
	passwordExpiryAt value_object.PasswordExpiryAt
}

func NewUser(
	name,
	email string,
	dateOfBirth value_object.DateOfBirth,
	password value_object.Password,
	expiryDuration time.Duration,
) *User {
	passwordExpiryAt := time.Now().Add(PasswordExpiryDuration)

	return &User{
		id:               shared_value_object.NewId(),
		name:             name,
		email:            email,
		dateOfBirth:      dateOfBirth,
		password:         password,
		passwordExpiryAt: value_object.NewPasswordExpiryAtFromTime(passwordExpiryAt),
	}
}

func (u *User) GetId() uuid.UUID {
	return u.id.GetId()
}

func (u *User) GetStringId() string {
	return u.id.GetStringId()
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) Rename(value string) error {
	if value == "" {
		return errors.New("name can't be empty")
	}

	u.name = value

	return nil
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(value string) {
	u.email = value
}

func (u *User) ResetPassword(value value_object.Password) {
	u.password = value
	// 重新设置完成后应该刷新过期时间
	passwordExpiryAt := time.Now().Add(PasswordExpiryDuration)
	u.passwordExpiryAt = value_object.NewPasswordExpiryAtFromTime(passwordExpiryAt)
}

func (u *User) Available() bool {
	return !u.passwordExpiryAt.IsExpired() && true
}
