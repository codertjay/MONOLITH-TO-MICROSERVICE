package products

import (
	"errors"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/price"
)

type Id string

var (
	ErrEmptyId   = errors.New("empty product id")
	ErrEmptyName = errors.New("empty product  name")
)

type Product struct {
	id          Id
	name        string
	description string
	price       price.Price
}

func NewProduct(
	id Id,
	name string,
	description string,
	price price.Price) (*Product, error) {
	if len(id) == 0 {
		return nil, ErrEmptyId
	}
	if len(name) == 0 {
		return nil, ErrEmptyName
	}
	return &Product{
		id,
		name,
		description,
		price}, nil

}

func (p Product) Id() Id {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() price.Price {
	return p.price
}
