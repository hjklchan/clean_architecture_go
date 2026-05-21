package main

import (
	"fmt"

	"practices.com/clean_arch_go/aggregate"
	"practices.com/clean_arch_go/infrastructure/repository/product"
)

func main() {
	repo := product.NewProductMemoryRepository()

	newProduct, err := aggregate.NewProduct("Lucas", "Lucas.Chen", 339.98)
	if err != nil {
		fmt.Println("new product err:", err)
		return
	}
	err = repo.Create(newProduct)
	if err != nil {
		fmt.Println("create product err:", err)
		return
	}

	products, err := repo.GetAll()
	if err != nil {
		fmt.Println("create product err:", err)
		return
	}

	for _, product := range products {
		fmt.Println(product.Display())
	}
}
