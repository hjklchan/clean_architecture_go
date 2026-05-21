package entity

import "github.com/google/uuid"

type Person struct {
	ID   uuid.UUID
	Name string
	Age  uint8
}

func NewPerson(name string, age uint8) *Person {
	return &Person{
		ID:   uuid.New(),
		Name: name,
		Age:  age,
	}
}

func DefaultPerson() *Person {
	return &Person{}
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() uint8 {
	return p.Age
}

func (p *Person) SetAge(age uint8) {
	p.Age = age
}
