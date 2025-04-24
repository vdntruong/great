package api

import (
	"encoding/json"
	"net/http"

	"oop/internal/domain"
	"oop/internal/services"
)

// ProductHandler handles HTTP requests for products
// Single Responsibility Principle: Only handles product-related HTTP requests
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProduct handles HTTP GET requests for a product
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from URL params
	id := r.URL.Query().Get("id")

	product, err := h.productService.GetProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// OrderHandler handles HTTP requests for orders
// Single Responsibility Principle: Only handles order-related HTTP requests
type OrderHandler struct {
	orderService *services.OrderService
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder handles HTTP POST requests to create an order
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var createOrderRequest struct {
		CustomerID string             `json:"customer_id"`
		Items      []domain.OrderItem `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&createOrderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.orderService.CreateOrder(
		r.Context(),
		createOrderRequest.CustomerID,
		createOrderRequest.Items,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// Router sets up HTTP routes
// Single Responsibility Principle: Only handles routing
type Router struct {
	productHandler *ProductHandler
	orderHandler   *OrderHandler
}

// NewRouter creates a new router
func NewRouter(productHandler *ProductHandler, orderHandler *OrderHandler) *Router {
	return &Router{
		productHandler: productHandler,
		orderHandler:   orderHandler,
	}
}

// SetupRoutes sets up HTTP routes
func (r *Router) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", r.productHandler.GetProduct)
	mux.HandleFunc("/orders", r.orderHandler.CreateOrder)

	return mux
}
