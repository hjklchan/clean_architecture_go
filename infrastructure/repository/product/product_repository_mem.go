package product

import (
	"sync"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/aggregate"
	"practices.com/clean_arch_go/domain/product"
)

type ProductMemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (repo *ProductMemoryRepository) GetAll() (products []aggregate.Product, err error) {
	for _, product := range repo.products {
		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductMemoryRepository) GetById(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := repo.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (repo *ProductMemoryRepository) Create(new aggregate.Product) error {
	repo.Lock()
	defer repo.Unlock()

	id := new.GetId()
	if _, ok := repo.products[id]; ok {
		return product.ErrProductAlreadyExist
	}
	repo.products[id] = new

	return nil
}

func (repo *ProductMemoryRepository) Update(new aggregate.Product) error {
	repo.Lock()
	defer repo.Unlock()

	_, ok := repo.products[new.GetId()]
	if !ok {
		return product.ErrProductNotFound
	}
	repo.products[new.GetId()] = new

	return nil
}

func (repo *ProductMemoryRepository) Delete(id uuid.UUID) error {
	repo.Lock()
	defer repo.Unlock()

	_, ok := repo.products[id]
	if !ok {
		return product.ErrProductNotFound
	}
	delete(repo.products, id)

	return nil
}
