package main

import (
	"fmt"
	"time"

	"practices.com/clean_arch_go/domain/user/value_object"
)

func main() {
	dob1 := value_object.NewDateOfBirth(time.Date(1999, 5, 19, 0, 0, 0, 0, time.Local))
	dob2 := value_object.NewDateOfBirth(time.Date(1999, 5, 20, 0, 0, 0, 0, time.Local))

	fmt.Println("age:", dob1.Age())
	fmt.Println("is adult:", dob1.IsAdult())

	fmt.Println("dob1 is younger than dob2?", dob1.IsYoungerThan(dob2))
	fmt.Println("dob1 is older than dob2?", dob1.IsOlderThan(dob2))
	fmt.Println("is birthday today?", dob2.IsBirthdayToday())
}
