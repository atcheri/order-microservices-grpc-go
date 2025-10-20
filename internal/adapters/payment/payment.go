package payment

import (
	"context"

	"github.com/atcheri/order-microservices-grpc-go/internal/application/domain"
	"github.com/atcheri/order-microservices-grpc-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := payment.NewPaymentClient(conn)

	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})

	return err
}
