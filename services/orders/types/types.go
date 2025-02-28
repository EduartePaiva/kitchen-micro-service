package types

import (
	"context"

	"github.com/eduartepaiva/kitchen-micro-service/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context, *orders.GetOrdersRequest) ([]*orders.Order, error)
}
