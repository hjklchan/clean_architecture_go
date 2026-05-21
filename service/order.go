package service

import (
	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/customer"
)

type OrderService struct {
	customerRepo customer.CustomerRepository
}

type OrderConfiguration func(srv *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	srv := &OrderService{}

	for _, cfg := range cfgs {
		if err := cfg(srv); err != nil {
			return nil, err
		}
	}

	return srv, nil
}

func (srv *OrderService) CreateOrder(
	customerId uuid.UUID,
	productIds []uuid.UUID,
) error {
	// Get customer from repo by id
	customer, err := srv.customerRepo.GetOne(customerId)
	if err != nil {
		return err
	}

	_ = customer
	// Get each product and need a product repo

	return nil
}
