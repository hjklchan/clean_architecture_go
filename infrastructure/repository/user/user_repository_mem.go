package user

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/user"
)

type UserMemoryRepository struct {
	users map[uuid.UUID]user.User
	sync.Mutex
}

func NewMockUserRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[uuid.UUID]user.User),
	}
}

func (repo *UserMemoryRepository) GetByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}

	return user.User{}, errors.New("the user was not found")
}

func (repo *UserMemoryRepository) Save(ctx context.Context, new user.User) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.users[new.GetId()]; ok {
		return errors.New("the user already exists")
	}

	repo.users[new.GetId()] = new

	return nil
}
