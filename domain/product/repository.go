package product

import (
	"errors"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/aggregate"
)

var (
	ErrProductNotFound     = errors.New("the product was not found")
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetById(uuid.UUID) (aggregate.Product, error)
	Create(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}

type ExtProductRepository interface {
	BatchCreate([]aggregate.Product) error
}
