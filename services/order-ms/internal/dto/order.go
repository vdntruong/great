package dto

import (
	"time"

	"order-ms/internal/models"

	"github.com/google/uuid"
)

type CreateOrderFromCartRequest struct {
	CartID string `json:"cart_id"`
}

type CreateOrderRequest struct {
	UserID uuid.UUID          `json:"user_id"`
	Items  []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

type UpdateOrderRequest struct {
	Status models.OrderStatus `json:"status"`
}

type OrderResponse struct {
	ID          uuid.UUID           `json:"id"`
	UserID      uuid.UUID           `json:"user_id"`
	Status      models.OrderStatus  `json:"status"`
	TotalAmount float64             `json:"total_amount"`
	Items       []OrderItemResponse `json:"items"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type OrderItemResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	Price     float64   `json:"price"`
	Subtotal  float64   `json:"subtotal"`
}

type ListOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}

func ToOrderResponse(order *models.Order) *OrderResponse {
	if order == nil {
		return nil
	}

	items := make([]OrderItemResponse, len(order.Items))
	for i, item := range order.Items {
		items[i] = OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
		}
	}

	return &OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Items:       items,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}

func ToOrdersResponse(orders []*models.Order) *ListOrdersResponse {
	if orders == nil {
		return &ListOrdersResponse{Orders: make([]OrderResponse, 0)}
	}

	responses := make([]OrderResponse, len(orders))
	for i, order := range orders {
		response := ToOrderResponse(order)
		if response != nil {
			responses[i] = *response
		}
	}

	return &ListOrdersResponse{Orders: responses}
}
