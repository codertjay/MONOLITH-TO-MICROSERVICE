package main

import (
	"fmt"
	"github.com/codertjay/MONOLITH-TO-MICROSERVICE/pkg/common/cmd"
	"log"
	"os"
)

func main() {
	log.Println("Starting payment microservice")
	defer log.Println("Closing payments microservice")
	ctx := cmd.Context()
	paymentsInterFace := createPaymentsMicroservice()
	if err := paymentsInterFace.Run(ctx); err != nil {
		panic(err)
	}

}

func createPaymentsMicroservice() amqp.PaymentsInterFace {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)
	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("ampq://%s", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHO_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)
	if err != nil {
		panic(err)
	}
	return paymentsInterface
}
