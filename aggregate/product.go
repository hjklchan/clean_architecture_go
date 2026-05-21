package aggregate

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"practices.com/clean_arch_go/entity"
)

// 产品创建时库存最少一件
const MIN_PRODUCT_QUANTITY = 1

type Amount struct {
	Intvalue   int64
	Floatvalue float64
}

func NewAmountFromInt(value int64) Amount {
	return Amount{
		Intvalue:   value,
		Floatvalue: float64(value / 100),
	}
}

func NewAmountFromFloat(value float64) Amount {
	return Amount{
		Intvalue:   int64(value * 100),
		Floatvalue: value,
	}
}

var (
	ErrEmptyName        = errors.New("field name can not be empty")
	ErrEmptyDescription = errors.New("field description can not be empty")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" {
		return Product{}, ErrEmptyName
	}
	if description == "" {
		return Product{}, ErrEmptyDescription
	}

	item := &entity.Item{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	return Product{
		item:     item,
		price:    price,
		quantity: MIN_PRODUCT_QUANTITY,
	}, nil
}

// 获取聚合根 Item 的 Id
func (p *Product) GetId() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) Display() string {
	return fmt.Sprintf("%s\t\t|%s\t\t|%f", p.GetItem().Name, p.GetItem().Description, p.GetPrice())
}
