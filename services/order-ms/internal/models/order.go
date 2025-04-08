package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipping  OrderStatus = "shipping"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCanceled  OrderStatus = "canceled"
)

type Order struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	Status      OrderStatus `json:"status"`
	TotalAmount float64     `json:"total_amount"`
	Items       []OrderItem `json:"items"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	Price     float64   `json:"price"`
	Subtotal  float64   `json:"subtotal"`
}

type CreateOrderParams struct {
	UserID uuid.UUID         `json:"user_id"`
	Items  []OrderItemParams `json:"items"`
}

type OrderItemParams struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

type UpdateOrderStatusParams struct {
	ID     uuid.UUID   `json:"id"`
	Status OrderStatus `json:"status"`
}

type ListOrdersParams struct {
	UserID   *uuid.UUID   `json:"user_id,omitempty"`
	Status   *OrderStatus `json:"status,omitempty"`
	Page     int32        `json:"page"`
	PageSize int32        `json:"page_size"`
}
