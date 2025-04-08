package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"order-ms/internal/dto"
	"order-ms/internal/models"
	"order-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Order struct {
	service   service.OrderService
	validator service.OrderValidator
}

func NewOrderHandler(orderService service.OrderService, validator service.OrderValidator) *Order {
	return &Order{
		service:   orderService,
		validator: validator,
	}
}

func (h *Order) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateOrder)
	r.Get("/{id}", h.GetOrder)
	r.Get("/", h.ListOrders)
	r.Put("/{id}/status", h.UpdateOrderStatus)
	r.Delete("/{id}", h.CancelOrder)

	return r
}

func (h *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert request to params
	params := models.CreateOrderParams{
		UserID: req.UserID,
		Items:  make([]models.OrderItemParams, len(req.Items)),
	}

	for i, item := range req.Items {
		params.Items[i] = models.OrderItemParams{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	order, err := h.service.CreateOrder(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ToOrderResponse(order))
}

func (h *Order) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(r.Context(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ToOrderResponse(order))
}

func (h *Order) ListOrders(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	page := 1
	pageSize := 10

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	// Parse status filter
	var status *models.OrderStatus
	if statusStr := r.URL.Query().Get("status"); statusStr != "" {
		s := models.OrderStatus(statusStr)
		status = &s
	}

	// List orders
	orders, err := h.service.ListOrders(r.Context(), models.ListOrdersParams{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Status:   status,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ToOrdersResponse(orders))
}

func (h *Order) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := models.UpdateOrderStatusParams{
		ID:     orderID,
		Status: req.Status,
	}

	order, err := h.service.UpdateOrderStatus(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.ToOrderResponse(order))
}

func (h *Order) CancelOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	err = h.service.CancelOrder(r.Context(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
