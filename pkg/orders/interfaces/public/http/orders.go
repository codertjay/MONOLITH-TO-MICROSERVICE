package http

import (
	common_http "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/http"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/orders/application"
	order "github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/orders/domain/orders"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func AddRoutes(router *chi.Mux, service application.OrderService, repository order.Repository) {
	resource := OrdersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

type PostOrderRequest struct {
	ProductId order.ProductId  `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}
type OrdersResource struct {
	service    application.OrderService
	repository order.Repository
}
type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}
type PostOrderResponse struct {
	OrderId string
}

type OrderPaidView struct {
	Id     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func (o OrdersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	o.repository.ById(order.Id(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
	}
	// fix the error in here
	render.Respond(w, r, OrderPaidView{string(order.Id()), order.Order.Paid()})
}

func (o OrdersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	render.Decode(r, &req)
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	cmd := application.PlaceOrderCommand{
		OrderId:   order.Id(uuid.NewV1().String()),
		ProductId: req.ProductId,
		Address:   application.PlaceOrderCommandAddress{req.Address},
	}
	if err = o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
	}
	w.WriteHeader(http.StatusOk)
	render.JSON(w, r, PostOrdersResponse{
		OrderId: string(cmd.OrderId),
	})
}
