package orders

import "errors"

type Address struct {
	name     string
	street   string
	city     string
	postCode string
	country  string
}

func NewAddress(
	name string,
	street string,
	city string,
	postCode string,
	country string) (Address, error) {

	if len(name) == 0 {
		return Address{}, errors.New("empty name")
	}
	if len(street) == 0 {
		return Address{}, errors.New("empty street")
	}
	if len(city) == 0 {
		return Address{}, errors.New("empty city")
	}
	if len(postCode) == 0 {
		return Address{}, errors.New("empty postCode")
	}
	if len(country) == 0 {
		return Address{}, errors.New("empty country")
	}
	return Address{
		name:     name,
		street:   street,
		city:     city,
		postCode: postCode,
		country:  country,
	}, nil
}
func (a Address) Street() {

}
func (a Address) Name() {

}
func (a Address) City() {

}
func (a Address) PostCode() {

}
func (a Address) Country() {

}
