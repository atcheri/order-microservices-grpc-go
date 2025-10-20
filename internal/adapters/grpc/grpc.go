package grpc

import (
	"context"

	"github.com/atcheri/order-microservices-grpc-go/internal/application/domain"
	"github.com/atcheri/order-microservices-grpc-proto/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderItems := make([]domain.OrderItem, len(request.OrderItems))

	for i, orderItem := range orderItems {
		orderItems[i] = domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		}
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)

	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		OrderId: result.ID,
	}, err
}
