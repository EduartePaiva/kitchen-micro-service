package service

import (
	"context"

	"github.com/eduartepaiva/kitchen-micro-service/services/common/genproto/orders"
)

var ordersDB = make([]*orders.Order, 0)

type ordersService struct {
	//store
}

func NewOrdersService() *ordersService {
	return &ordersService{}
}

func (s *ordersService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

func (s *ordersService) GetOrders(
	ctx context.Context,
	orderID *orders.GetOrdersRequest,
) ([]*orders.Order, error) {
	userOrders := make([]*orders.Order, 0)
	CustomerID := orderID.CustomerID
	for _, order := range ordersDB {
		if CustomerID == order.CustomerID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders, nil
}
