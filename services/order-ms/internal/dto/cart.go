package dto

import (
	"time"

	"order-ms/internal/models"

	"github.com/google/uuid"
)

type AddToCartRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

type UpdateCartItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

type CartItemResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
}

type CartResponse struct {
	ID        uuid.UUID          `json:"id"`
	UserID    uuid.UUID          `json:"user_id"`
	Items     []CartItemResponse `json:"items"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func ToCartResponse(cart *models.Cart) *CartResponse {
	if cart == nil {
		return nil
	}

	items := make([]CartItemResponse, len(cart.Items))
	for i, item := range cart.Items {
		items[i] = CartItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	return &CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		Items:     items,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}
