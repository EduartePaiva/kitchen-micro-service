package main

import (
	"log"
	"net"

	handler "github.com/eduartepaiva/kitchen-micro-service/services/orders/handler/orders"
	"github.com/eduartepaiva/kitchen-micro-service/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) Run() error {
	httpListener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	grpcServer := grpc.NewServer()
	// register our grpc services
	orderService := service.NewOrdersService()
	handler.NewGrpcOrdersService(grpcServer, orderService)
	log.Println("grpc service is listening on ", s.addr)
	return grpcServer.Serve(httpListener)
}
