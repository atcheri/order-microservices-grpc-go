package ports

import "github.com/atcheri/order-microservices-grpc-go/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
