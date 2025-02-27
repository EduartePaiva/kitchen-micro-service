package handler

import (
	"context"

	"github.com/eduartepaiva/kitchen-micro-service/services/common/genproto/orders"
	"github.com/eduartepaiva/kitchen-micro-service/services/orders/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdersGrpcHandler struct {
	// service injection
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{ordersService: ordersService}
	// register the orderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}
	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &orders.CreateOrderResponse{Status: "success"}, nil
}
func (OrdersGrpcHandler) GetOrders(context.Context, *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
