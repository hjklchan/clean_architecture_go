package customer

import (
	"errors"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("repo: the customer was not found")
	ErrCustomerFailedToAdd    = errors.New("repo: failed to add the customer")
	ErrCustomerFailedToUpdate = errors.New("repo: failed to update the customer")
)

type CustomerRepository interface {
	GetOne(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
