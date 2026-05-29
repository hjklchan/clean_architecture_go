package main

import (
	"fmt"

	"practices.com/clean_arch_go/value_object/money"
)

func main() {
	price1, err := money.NewMoney("rmb", 1.02)
	price2, err := money.NewMoney("rmb", 5.97)
	if err != nil {
		fmt.Println(err)
		return
	}

	gap, err := price1.Gap(price2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(gap)
}
