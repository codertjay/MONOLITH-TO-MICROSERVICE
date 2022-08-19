package products

import "errors"

/*
we use this file to talk to our memory.go file that
i what repository is meant for .
The repository is just an interface

*/

var ErrNotFound = errors.New("product not found")

type Repository interface {
	Save(*Product) error
	ById(Id) (*Product, error)
}
