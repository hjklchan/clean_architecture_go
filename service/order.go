package service

import (
	"github.com/google/uuid"
	"practices.com/clean_arch_go/domain/customer"
	"practices.com/clean_arch_go/domain/product"
)

type OrderService struct {
	customerRepo customer.CustomerRepository
	productRepo  product.ProductRepository
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
) (float64, error) {
	// Get customer from repo by id
	customer, err := srv.customerRepo.GetOne(customerId)
	if err != nil {
		return 0, err
	}

	_ = customer
	var totalPrice float64
	// Get each product and need a product repo
	for _, pid := range productIds {
		p, err := srv.productRepo.GetById(pid)
		if err != nil {
			return 0, err
		}
		totalPrice += p.GetPrice()
	}

	return totalPrice, nil
}
