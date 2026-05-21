package customer

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/aggregate"
	"practices.com/clean_arch_go/domain/customer"
)

type CustomerMemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *CustomerMemoryRepository {
	return &CustomerMemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (repo *CustomerMemoryRepository) GetOne(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := repo.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *CustomerMemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrCustomerFailedToAdd)
	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}

func (mr *CustomerMemoryRepository) Update(c aggregate.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrCustomerFailedToUpdate)
	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}
