package http

import (
	common_http "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/http"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/price"
	products "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/shop/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type ProductsResource struct {
	readModel ProductsReadModel
}

func AddRoutes(router *chi.Mux, productsReadModel ProductsReadModel) {
	resource := ProductsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type ProductsReadModel interface {
	AllProducts() ([]products.Product, error)
}

type ProductView struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       PriceView `json:"price"`
}
type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func PriceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func (p ProductsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	allProducts, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	view := []ProductView{}
	for _, product := range allProducts {
		view = append(view, ProductView{
			string(product.Id()),
			product.Name(),
			product.Description(),
			PriceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
