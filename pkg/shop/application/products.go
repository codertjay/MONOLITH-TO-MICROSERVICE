package application

import (
	"errors"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/price"
	products "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/shop/domain"
)

type ProductReadModel interface {
	AllProducts() ([]products.Product, err)
}
type ProductService struct {
	repo      products.Repository
	readModel ProductReadModel
}

func NewProductService(
	repo products.Repository,
	readModel ProductReadModel) ProductService {
	return ProductService{repo: repo, readModel: readModel}
}

func (s ProductService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

type AddProductCommand struct {
	Id            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductService) AddProducts(cmd AddProductCommand) error {
	newPrice, err := price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "invalid product price")
	}
	product, err := products.NewProduct(products.Id(cmd.Id), cmd.Name, cmd.Description, newPrice)
	if err != nil {
		return errors.Wrap(err, "cannot create product")
	}

	err = s.repo.Save(product)
	if err != nil {
		return errors.Wrap(err, "cannot save product")
	}
	return nil
}
