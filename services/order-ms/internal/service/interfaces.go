package service

import (
	"context"

	"order-ms/internal/models"

	"github.com/google/uuid"
)

type CartService interface {
	CreateCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error)
	GetCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error)
	AddItem(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error
	UpdateItem(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error
	RemoveItem(ctx context.Context, cartID uuid.UUID, productID uuid.UUID) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, params models.CreateOrderParams) (*models.Order, error)
	GetOrder(ctx context.Context, id uuid.UUID) (*models.Order, error)
	ListOrders(ctx context.Context, params models.ListOrdersParams) ([]*models.Order, error)
	UpdateOrderStatus(ctx context.Context, params models.UpdateOrderStatusParams) (*models.Order, error)
	CancelOrder(ctx context.Context, id uuid.UUID) error
}
