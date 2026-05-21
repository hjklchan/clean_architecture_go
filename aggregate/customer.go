package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
	"practices.com/clean_arch_go/value_object"
)

// 聚合类
type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []value_object.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, errors.New("name cannot be empty for customer")
	}

	customer := Customer{
		person: &entity.Person{
			Name: name,
			ID:   uuid.New(),
		},
		products:     make([]*entity.Item, 0),
		transactions: make([]value_object.Transaction, 0),
	}

	return customer, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = entity.DefaultPerson()
	}

	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.GetName()
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = entity.DefaultPerson()
	}

	c.person.SetName(name)
}
