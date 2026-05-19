package value_object

import (
	"time"
)

type DateOfBirth struct {
	value time.Time
}

func NewDateOfBirth(value time.Time) DateOfBirth {
	return DateOfBirth{
		value,
	}
}

func (dob DateOfBirth) Age() int {
	y, m, d := time.Now().Date()
	age := y - dob.value.Year()

	if m >= dob.value.Month() && d >= dob.value.Day() {
		return age
	} else {
		return age - 1
	}
}

func (dob DateOfBirth) IsAdult() bool {
	return dob.Age() > 18
}

func (dob DateOfBirth) IsYoungerThan(other DateOfBirth) bool {
	return dob.value.After(other.value)
}

func (dob DateOfBirth) IsOlderThan(other DateOfBirth) bool {
	return !dob.IsYoungerThan(other)
}

func (dob DateOfBirth) IsBirthdayToday() bool {
	_, m, d := time.Now().Date()

	return dob.value.Month() == m && dob.value.Day() == d
}

func (dob DateOfBirth) Value() time.Time {
	return dob.value
}
