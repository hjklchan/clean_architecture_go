package main

import (
	"fmt"
	"time"

	"practices.com/clean_arch_go/value_object/date_range"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	fromObj1, err := time.ParseInLocation("2006/01/02 15:04:05", "2026/06/01 00:00:00", loc)
	toObj1, err := time.ParseInLocation("2006/01/02 15:04:05", "2026/06/02 10:00:00", loc)

	dateRange, err := date_range.NewDateRange(fromObj1, toObj1)
	if err != nil {
		fmt.Println("date range err:", err)
		return
	}

	fmt.Println("date range obj:", dateRange.Difference().Hours())
}
