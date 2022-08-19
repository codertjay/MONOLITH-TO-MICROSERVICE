package main

import (
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/cmd"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting the orders micro service")
	ctx := cmd.Context()

	r, closeFn := createOrderMicroService()
	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	// check this code if it supposes to be inside the go func
	<-ctx.Done()
	log.Println("Closing order microservice")
	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createOrderMicroService() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP+RABBITMQ_ADDR"))
	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.CreateRouter()
	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)
	return r, func() {

	}
}
