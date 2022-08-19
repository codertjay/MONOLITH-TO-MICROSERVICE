package http

import (
	common_http "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/http"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/price"
	products_domain "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/shop/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := ProductsResource{repo: repo}
	router.Get("/product/{id}", resource.Get)
}

type ProductsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

type ProductView struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       PriceView `json:"price"`
}

func PriceViewFromPrice(p price.Price) PriceView {
	return PriceView{p.Cents(), p.Currency()}
}

func (p ProductsResource) Get(w http.ResponseWriter, r *http.Request) {
	current_product, err := p.repo.ById(products_domain.Id(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	// todo : fix bug
	// fixme: fix bug he is using PriceView as the struct but wil lke to use product struct
	render.Respond(w, r, ProductView{
		string(current_product.Id()),
		current_product.Name(),
		current_product.Description(),
		PriceViewFromPrice(current_product.Price()),
	})
}
