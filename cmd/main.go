package main

import (
	"log"

	"github.com/atcheri/order-microservices-grpc-go/config"
	"github.com/atcheri/order-microservices-grpc-go/internal/adapters/db"
	"github.com/atcheri/order-microservices-grpc-go/internal/adapters/grpc"
	"github.com/atcheri/order-microservices-grpc-go/internal/adapters/payment"
	"github.com/atcheri/order-microservices-grpc-go/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())

	if err != nil {

		log.Fatalf("failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
