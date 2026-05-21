package main

import (
	"fmt"

	"practices.com/clean_arch_go/aggregate"
)

func main() {
	customer1, err := aggregate.NewCustomer("Lucas")
	if err != nil {
		fmt.Println("created customer1 err:", err)
		return
	}
	fmt.Printf("%#v\n", customer1)

	customer2, err := aggregate.NewCustomer("")
	if err != nil {
		fmt.Println("created customer2 err:", err)
		return
	}
	fmt.Printf("%#v\n", customer2)
}
