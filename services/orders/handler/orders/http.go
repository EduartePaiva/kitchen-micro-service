package handler

import (
	"net/http"

	"github.com/eduartepaiva/kitchen-micro-service/services/common/genproto/orders"
	"github.com/eduartepaiva/kitchen-micro-service/services/common/utils"
	"github.com/eduartepaiva/kitchen-micro-service/services/orders/types"
)

type ordersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(ordersService types.OrderService) *ordersHttpHandler {
	return &ordersHttpHandler{ordersService: ordersService}
}

func (h *ordersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /api/orders", h.CreateOrder)
	router.HandleFunc("GET /api/orders", h.GetOrders)
}

func (h ordersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var reqOrder orders.CreateOrderRequest
	err := utils.ParseJSON(r, &reqOrder)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	order := &orders.Order{
		OrderID:    42,
		CustomerID: reqOrder.CustomerID,
		ProductID:  reqOrder.ProductID,
		Quantity:   reqOrder.Quantity,
	}
	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	utils.WriteJSON(w, http.StatusOK, res)
}
func (ordersHttpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
}
