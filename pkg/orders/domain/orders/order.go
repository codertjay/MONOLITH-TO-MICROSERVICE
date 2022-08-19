package orders

import (
	"errors"
)

type Id string

var ErrEmptyOrderId = errors.New("empty order id")

type Order struct {
	id      Id
	product Product
	address Address
	paid    bool
}

func (o *Order) Id() Id {
	return o.Id()
}
func (o Order) Product() Product {
	return o.product
}

func (o *Order) Address() Address {
	return o.address
}

func (o *Order) Paid() bool {
	return o.paid
}
func (o *Order) MarkAsPaid() {
	o.paid = true
}

func NewOrder(id Id, product Product, address Address) (*Order, error) {
	if len(id) == 0 {
		return nil, ErrEmptyOrderId
	}
	return &Order{id, product, address, false}, nil
}
