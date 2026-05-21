package value_object

type Address struct {
	Country string
	State   string
	City    string
	Street  string
	ZipCode string
}

func NewAddress(country, state, city, street, zipCode string) Address {
	return Address{
		Country: country,
		State:   state,
		City:    city,
		Street:  street,
		ZipCode: zipCode,
	}
}

func (a Address) EqualTo(target Address) bool {
	return a.Country == target.Country &&
		a.State == target.State &&
		a.City == target.City &&
		a.Street == target.Street &&
		a.ZipCode == target.ZipCode
}
