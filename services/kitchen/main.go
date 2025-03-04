package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error creating grpc client:", err)
	}
	return conn
}

func main() {
	sv := NewHttpServer(":3000")
	log.Fatal(sv.Run())
}
