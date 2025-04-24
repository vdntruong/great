package domain

import (
	"errors"
	"time"
)

// Product represents a product in our e-commerce system
type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Inventory   int
}

// ValidateProduct checks if a product is valid
func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if p.Price <= 0 {
		return errors.New("product price must be positive")
	}
	if p.Inventory < 0 {
		return errors.New("product inventory cannot be negative")
	}
	return nil
}

// Order represents a customer order
type Order struct {
	ID         string
	CustomerID string
	Items      []OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
}

// OrderItem represents a product in an order
type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// CalculateTotal calculates the total price of an order
func (o *Order) CalculateTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
