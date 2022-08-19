package orders

import (
	"errors"
	price "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/price"
)

type ProductId string

var ErrorEmptyProductId = errors.New("empty product id")

type Product struct {
	id    ProductId
	name  string
	price price.Price
}

func (p Product) Id() ProductId {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}

func NewProduct(id ProductId, name string, price price.Price) (Product, error) {
	if len(name) == 0 {
		return Product{}, ErrorEmptyProductId
	}
	return Product{id, name, price}, nil
}
