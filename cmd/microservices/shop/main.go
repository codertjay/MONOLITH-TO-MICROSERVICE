package main

import (
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/cmd"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting shop microservice")

	ctx := cmd.Context()

	r := createShopMicroService()
	server := &http.Server{Addr: os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-ctx.Done()
	log.Println("Closing shop microservice")
	if err := server.Close(); err != nil {
		panic(err)
	}

}

func createShopMicroService() *chi.Mux {
	shopProductRepo := shop_infra_product.NewMemoryRepository()
	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)
	return r
}
