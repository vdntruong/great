package repositories

import (
	"context"

	"oop/internal/domain"
)

// ProductRepository defines operations for product persistence
// Single Responsibility Principle: Only responsible for product persistence
type ProductRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, error)
	Save(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id string) error
}

// OrderRepository defines operations for order persistence
// Single Responsibility Principle: Only responsible for order persistence
type OrderRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Order, error)
	FindByCustomer(ctx context.Context, customerID string) ([]*domain.Order, error)
	Save(ctx context.Context, order *domain.Order) error
	Update(ctx context.Context, order *domain.Order) error
}
