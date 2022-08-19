package products

import products "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/shop/domain"

type MemoryRepository struct {
	products []products.Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]products.Product{}}
}

func (m *MemoryRepository) Save(
	productToSave *products.Product) error {
	for i, p := range m.products {
		if p.Id() == productToSave.Id() {
			m.products[i] = *productToSave
			return nil
		}
	}
	m.products = append(m.products, *productToSave)
	return nil
}

func (m MemoryRepository) ById(id products.Id) (*products.Product, error) {
	for _, p := range m.products {
		if p.Id() == id {
			return &p, nil
		}
	}
	return nil, products.ErrNotFound
}

func (m MemoryRepository) AllProducts() ([]products.Product, error) {
	return m.products, nil
}
