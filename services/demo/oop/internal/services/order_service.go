package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"oop/internal/domain"
	"oop/internal/repositories"
)

// OrderService provides business logic for orders
// Single Responsibility Principle: Only handles order-related business logic
type OrderService struct {
	orderRepo   repositories.OrderRepository
	productRepo repositories.ProductRepository
	// Open/Closed Principle: We can add new dependencies without modifying existing code
	notifier OrderNotifier
}

// OrderNotifier defines the notification behavior
// Interface Segregation Principle: Small, focused interface
type OrderNotifier interface {
	NotifyOrderCreated(order *domain.Order) error
	NotifyOrderStatusChanged(order *domain.Order, oldStatus domain.OrderStatus) error
}

// NewOrderService creates a new order service
func NewOrderService(
	orderRepo repositories.OrderRepository,
	productRepo repositories.ProductRepository,
	notifier OrderNotifier,
) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		notifier:    notifier,
	}
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(ctx context.Context, customerID string, items []domain.OrderItem) (*domain.Order, error) {
	// Check if all products exist and have sufficient inventory
	for i, item := range items {
		product, err := s.productRepo.FindByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}
		if product == nil {
			return nil, errors.New("product not found")
		}

		if product.Inventory < item.Quantity {
			return nil, errors.New("insufficient inventory")
		}

		// Set the correct price from the product
		items[i].Price = product.Price
	}

	// Create the order
	order := &domain.Order{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		Items:      items,
		Status:     domain.OrderStatusPending,
		CreatedAt:  time.Now(),
	}

	// Save the order
	if err := s.orderRepo.Save(ctx, order); err != nil {
		return nil, err
	}

	// Update inventory for each product
	for _, item := range items {
		if err := s.productRepo.FindByID(ctx, item.ProductID); err != nil {
			// This is a critical error since we've already created the order
			// In a real system, we'd need proper transaction handling
			return order, errors.New("failed to update inventory after order creation")
		}
	}

	// Notify about order creation
	if err := s.notifier.NotifyOrderCreated(order); err != nil {
		// Log the error but don't fail the order creation
		// In a real system, we might use a retry mechanism
	}

	return order, nil
}

// UpdateOrderStatus updates an order's status
func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID string, newStatus domain.OrderStatus) error {
	order, err := s.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("order not found")
	}

	oldStatus := order.Status
	order.Status = newStatus

	if err := s.orderRepo.Update(ctx, order); err != nil {
		return err
	}

	// Notify about status change
	if err := s.notifier.NotifyOrderStatusChanged(order, oldStatus); err != nil {
		// Log the error but don't fail the status update
	}

	return nil
}
