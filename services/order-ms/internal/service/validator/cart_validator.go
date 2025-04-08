package validator

import (
	"fmt"
	"order-ms/internal/models"

	"github.com/google/uuid"
)

type CartValidator struct{}

func NewCartValidator() *CartValidator {
	return &CartValidator{}
}

func (v *CartValidator) ValidateGetCart(userID uuid.UUID) error {
	if userID == uuid.Nil {
		return fmt.Errorf("user ID is required")
	}
	return nil
}

func (v *CartValidator) ValidateAddItem(userID uuid.UUID, item *models.CartItem) error {
	if userID == uuid.Nil {
		return fmt.Errorf("user ID is required")
	}
	if item == nil {
		return fmt.Errorf("item is required")
	}
	if item.ProductID == uuid.Nil {
		return fmt.Errorf("product ID is required")
	}
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	return nil
}

func (v *CartValidator) ValidateUpdateItem(cartID uuid.UUID, item *models.CartItem) error {
	if cartID == uuid.Nil {
		return fmt.Errorf("cart ID is required")
	}
	if item == nil {
		return fmt.Errorf("item is required")
	}
	if item.ProductID == uuid.Nil {
		return fmt.Errorf("product ID is required")
	}
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	return nil
}

func (v *CartValidator) ValidateClearCart(cartID uuid.UUID) error {
	if cartID == uuid.Nil {
		return fmt.Errorf("cart ID is required")
	}
	return nil
}

func (v *CartValidator) ValidateRemoveItem(cartID uuid.UUID, productID uuid.UUID) error {
	if cartID == uuid.Nil {
		return fmt.Errorf("cart ID is required")
	}
	if productID == uuid.Nil {
		return fmt.Errorf("product ID is required")
	}
	return nil
}
