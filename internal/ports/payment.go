package ports

import (
	"github.com/atcheri/order-microservices-grpc-go/internal/application/domain"
)

type PaymentPort interface {
	Charge(*domain.Order) error
}
