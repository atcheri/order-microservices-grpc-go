package ports

import "github.com/atcheri/order-microservices-grpc-go/internal/application/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
