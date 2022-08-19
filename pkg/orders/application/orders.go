package application

type ProductService interface{}

type PaymentService interface {
}

type OrderService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommand struct {
}

func (*OrderService) PlaceOrderService(cmd PlaceOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
}

func (s OrderService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {

}
func (s OrderService) OrderById(id orders.Id) (orders.Order, error) {

}
