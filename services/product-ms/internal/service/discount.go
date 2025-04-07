package service

import (
	"context"
	"fmt"

	"product-ms/db/dao"
	"product-ms/internal/models"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// DiscountServiceImpl implements DiscountService
type DiscountServiceImpl struct {
	queries   *dao.Queries
	validator validator.DiscountValidator
}

var _ DiscountService = (*DiscountServiceImpl)(nil)

// NewDiscountService creates a new DiscountService
func NewDiscountService(queries *dao.Queries) *DiscountServiceImpl {
	return &DiscountServiceImpl{
		queries:   queries,
		validator: validator.NewDiscountValidator(),
	}
}

// CreateDiscount creates a new discount
func (s *DiscountServiceImpl) CreateDiscount(ctx context.Context, params models.CreateDiscountParams) (*models.Discount, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertCreateDiscountParamsToDAO(params)

	discount, err := s.queries.CreateDiscount(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertDiscountToModel(discount), nil
}

// GetDiscountByID gets a discount by ID
func (s *DiscountServiceImpl) GetDiscountByID(ctx context.Context, id string) (*models.Discount, error) {
	discountID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	discount, err := s.queries.GetDiscount(ctx, discountID)
	if err != nil {
		return nil, err
	}

	return ConvertDiscountToModel(discount), nil
}

// ListDiscounts lists discounts
func (s *DiscountServiceImpl) ListDiscounts(ctx context.Context, params models.ListDiscountsParams) (*models.DiscountList, error) {
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	daoParams := ConvertListDiscountsParamsToDAO(params)
	discounts, err := s.queries.ListDiscounts(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	total, err := s.queries.CountDiscounts(ctx, params.StoreID)
	if err != nil {
		return nil, err
	}

	return ConvertDiscountListToModel(discounts, total, int(params.Offset/params.Limit+1), int(params.Limit)), nil
}

// UpdateDiscount updates a discount
func (s *DiscountServiceImpl) UpdateDiscount(ctx context.Context, params models.UpdateDiscountParams) (*models.Discount, error) {
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertUpdateDiscountParamsToDAO(params)
	discount, err := s.queries.UpdateDiscount(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertDiscountToModel(discount), nil
}

// DeleteDiscount deletes a discount
func (s *DiscountServiceImpl) DeleteDiscount(ctx context.Context, id string) error {
	discountID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.queries.DeleteDiscount(ctx, discountID)
}

// AddDiscountProduct adds a product to a discount
func (s *DiscountServiceImpl) AddDiscountProduct(ctx context.Context, discountID, productID uuid.UUID) error {
	params := &dao.AddDiscountProductParams{
		DiscountID: discountID,
		ProductID:  productID,
	}

	if err := s.queries.AddDiscountProduct(ctx, params); err != nil {
		return fmt.Errorf("failed to add product to discount: %w", err)
	}

	return nil
}

// RemoveDiscountProduct removes a product from a discount
func (s *DiscountServiceImpl) RemoveDiscountProduct(ctx context.Context, discountID, productID uuid.UUID) error {
	params := &dao.RemoveDiscountProductParams{
		DiscountID: discountID,
		ProductID:  productID,
	}

	if err := s.queries.RemoveDiscountProduct(ctx, params); err != nil {
		return fmt.Errorf("failed to remove product from discount: %w", err)
	}

	return nil
}

// AddDiscountCategory adds a category to a discount
func (s *DiscountServiceImpl) AddDiscountCategory(ctx context.Context, discountID, categoryID uuid.UUID) error {
	params := &dao.AddDiscountCategoryParams{
		DiscountID: discountID,
		CategoryID: categoryID,
	}

	if err := s.queries.AddDiscountCategory(ctx, params); err != nil {
		return fmt.Errorf("failed to add category to discount: %w", err)
	}

	return nil
}

// RemoveDiscountCategory removes a category from a discount
func (s *DiscountServiceImpl) RemoveDiscountCategory(ctx context.Context, discountID, categoryID uuid.UUID) error {
	params := &dao.RemoveDiscountCategoryParams{
		DiscountID: discountID,
		CategoryID: categoryID,
	}

	if err := s.queries.RemoveDiscountCategory(ctx, params); err != nil {
		return fmt.Errorf("failed to remove category from discount: %w", err)
	}

	return nil
}
