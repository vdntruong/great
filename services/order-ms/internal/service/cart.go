package service

import (
	"context"
	"fmt"

	"order-ms/db/dao"
	"order-ms/internal/models"
	"order-ms/internal/service/validator"

	"github.com/google/uuid"
)

type CartValidator interface {
	ValidateGetCart(userID uuid.UUID) error
	ValidateAddItem(userID uuid.UUID, item *models.CartItem) error
	ValidateUpdateItem(cartID uuid.UUID, item *models.CartItem) error
	ValidateRemoveItem(cartID uuid.UUID, productID uuid.UUID) error
	ValidateClearCart(cartID uuid.UUID) error
}

type CartDAOQuerier interface {
	CreateCart(ctx context.Context, userID uuid.UUID) (dao.Cart, error)
	AddCartItem(ctx context.Context, arg dao.AddCartItemParams) (dao.CartItem, error)

	GetCart(ctx context.Context, userID uuid.UUID) (dao.GetCartRow, error)
	DeleteCart(ctx context.Context, id uuid.UUID) error

	RemoveCartItem(ctx context.Context, arg dao.RemoveCartItemParams) error
	UpdateCartItem(ctx context.Context, arg dao.UpdateCartItemParams) (dao.CartItem, error)
	ClearCart(ctx context.Context, cartID uuid.UUID) error
}

type CartServiceAdapter struct {
	querier   CartDAOQuerier
	validator CartValidator
}

var _ CartService = (*CartServiceAdapter)(nil)

func NewCartServiceAdapter(querier CartDAOQuerier) *CartServiceAdapter {
	return &CartServiceAdapter{
		querier:   querier,
		validator: validator.NewCartValidator(),
	}
}

func (s *CartServiceAdapter) CreateCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	_, err := s.querier.CreateCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create cart: %w", err)
	}

	cartRow, err := s.querier.GetCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	return s.convertCartToModel(ctx, cartRow), nil
}

func (s *CartServiceAdapter) GetCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	cartRow, err := s.querier.GetCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	return s.convertCartToModel(ctx, cartRow), nil
}

func (s *CartServiceAdapter) AddItem(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error {
	// Get cart
	cart, err := s.querier.GetCart(ctx, cartID)
	if err != nil {
		return err
	}

	// Validate params
	if err := s.validator.ValidateAddItem(cart.UserID, item); err != nil {
		return err
	}

	// Add item to cart
	_, err = s.querier.AddCartItem(ctx, dao.AddCartItemParams{
		CartID:    cartID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	})
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}

	return nil
}

func (s *CartServiceAdapter) RemoveItem(ctx context.Context, cartID uuid.UUID, productID uuid.UUID) error {
	// Validate params
	if err := s.validator.ValidateRemoveItem(cartID, productID); err != nil {
		return err
	}

	// Remove item
	return s.querier.RemoveCartItem(ctx, dao.RemoveCartItemParams{
		CartID:    cartID,
		ProductID: productID,
	})
}

func (s *CartServiceAdapter) UpdateItem(ctx context.Context, cartID uuid.UUID, item *models.CartItem) error {
	// Validate params
	if err := s.validator.ValidateUpdateItem(cartID, item); err != nil {
		return err
	}

	// Update item
	_, err := s.querier.UpdateCartItem(ctx, dao.UpdateCartItemParams{
		CartID:    cartID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	})
	if err != nil {
		return fmt.Errorf("failed to update item in cart: %w", err)
	}
	return nil
}

func (s *CartServiceAdapter) convertCartToModel(ctx context.Context, cart dao.GetCartRow) *models.Cart {
	return &models.Cart{
		ID:     cart.ID,
		UserID: cart.UserID,
		Items:  make([]models.CartItem, 0),
	}
}
