package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/eduartepaiva/kitchen-micro-service/services/common/genproto/orders"
	"github.com/eduartepaiva/kitchen-micro-service/services/common/utils"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()
	conn := NewGRPCClient(":9000")
	defer conn.Close()
	ordersClient := orders.NewOrderServiceClient(conn)
	router.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()
		res, err := ordersClient.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 2,
			ProductID:  123,
			Quantity:   2,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, res)
	})

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		orderRequest := &orders.GetOrdersRequest{}
		err := utils.ParseJSON(r, orderRequest)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()
		cltOrders, err := ordersClient.GetOrders(ctx, orderRequest)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, cltOrders)
	})
	log.Println("starting kitchen server on", h.addr)
	return http.ListenAndServe(h.addr, router)
}
