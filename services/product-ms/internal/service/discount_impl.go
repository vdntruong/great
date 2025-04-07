package service

import (
	"context"
	"database/sql"
	"fmt"

	"product-ms/internal/models"
	"product-ms/internal/repository/dao"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// DiscountServiceImpl implements DiscountService
type DiscountServiceImpl struct {
	dao       *dao.Queries
	validator validator.DiscountValidator
}

var _ DiscountService = (*DiscountServiceImpl)(nil)

// NewDiscountService creates a new DiscountService
func NewDiscountService(dao *dao.Queries) *DiscountServiceImpl {
	return &DiscountServiceImpl{
		dao:       dao,
		validator: validator.NewDiscountValidator(),
	}
}

// ToDiscountModel converts a DAO discount to a model discount
func ToDiscountModel(d *dao.Discount) *models.Discount {
	if d == nil {
		return nil
	}

	return &models.Discount{
		ID:                d.ID,
		StoreID:           d.StoreID,
		Name:              d.Name,
		Code:              d.Code,
		Type:              models.DiscountType(d.Type),
		Value:             stringToFloat64(d.Value),
		Scope:             models.DiscountScope(d.Scope),
		StartDate:         d.StartDate,
		EndDate:           nullTimeToTimePtr(d.EndDate),
		MinPurchaseAmount: stringToFloat64Ptr(d.MinPurchaseAmount),
		MaxDiscountAmount: stringToFloat64Ptr(d.MaxDiscountAmount),
		UsageLimit:        nullInt32ToInt32Ptr(d.UsageLimit),
		UsageCount:        nullInt32ToInt32(d.UsageCount),
		IsActive:          d.IsActive.Bool,
		CreatedAt:         d.CreatedAt.Time,
		UpdatedAt:         d.UpdatedAt.Time,
	}
}

// ToDiscountListModel converts a DAO discount list to a model discount list
func ToDiscountListModel(discounts []*dao.Discount, total int64) *models.DiscountList {
	items := make([]models.Discount, len(discounts))
	for i, d := range discounts {
		items[i] = *ToDiscountModel(d)
	}

	return &models.DiscountList{
		Discounts:   items,
		TotalCount: total,
		Page:       1, // TODO: Calculate page based on offset and limit
		Limit:      10, // TODO: Use actual limit
	}
}

// CreateDiscount creates a new discount
func (s *DiscountServiceImpl) CreateDiscount(ctx context.Context, params models.CreateDiscountParams) (*models.Discount, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	id := uuid.New()
	daoParams := dao.CreateDiscountParams{
		ID:                id,
		StoreID:           params.StoreID,
		Name:              params.Name,
		Code:              params.Code,
		Type:              dao.DiscountType(params.Type),
		Value:             float64ToString(params.Value),
		Scope:             dao.DiscountScope(params.Scope),
		StartDate:         params.StartDate,
		EndDate:           timePtrToNullTime(params.EndDate),
		MinPurchaseAmount: float64PtrToString(params.MinPurchaseAmount),
		MaxDiscountAmount: float64PtrToString(params.MaxDiscountAmount),
		UsageLimit:        int32PtrToNullInt32(params.UsageLimit),
		IsActive:          sql.NullBool{Bool: params.IsActive, Valid: true},
	}

	discount, err := s.dao.CreateDiscount(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToDiscountModel(discount), nil
}

// GetDiscountByID gets a discount by ID
func (s *DiscountServiceImpl) GetDiscountByID(ctx context.Context, id string) (*models.Discount, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	discount, err := s.dao.GetDiscountByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return ToDiscountModel(discount), nil
}

// ListDiscounts lists discounts
func (s *DiscountServiceImpl) ListDiscounts(ctx context.Context, params models.ListDiscountsParams) (*models.DiscountList, error) {
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	daoParams := dao.ListDiscountsParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}

	discounts, err := s.dao.ListDiscounts(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToDiscountListModel(discounts, int64(len(discounts))), nil
}

// UpdateDiscount updates a discount
func (s *DiscountServiceImpl) UpdateDiscount(ctx context.Context, params models.UpdateDiscountParams) (*models.Discount, error) {
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	daoParams := dao.UpdateDiscountParams{
		ID:         params.ID,
		Column2:    params.Name,
		Column3:    params.Code,
		Column4:    (*string)(params.Type),
		Column5:    float64PtrToString(params.Value),
		Column6:    (*string)(params.Scope),
		StartDate:  *params.StartDate,
		EndDate:    timePtrToNullTime(params.EndDate),
		Column9:    float64PtrToString(params.MinPurchaseAmount),
		Column10:   float64PtrToString(params.MaxDiscountAmount),
		UsageLimit: int32PtrToNullInt32(params.UsageLimit),
		IsActive:   boolPtrToNullBool(params.IsActive),
	}

	discount, err := s.dao.UpdateDiscount(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToDiscountModel(discount), nil
}

// DeleteDiscount deletes a discount
func (s *DiscountServiceImpl) DeleteDiscount(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.dao.DeleteDiscount(ctx, uuid)
}

// AddDiscountProduct adds a product to a discount
func (s *DiscountServiceImpl) AddDiscountProduct(ctx context.Context, discountID, productID uuid.UUID) error {
	params := &dao.AddDiscountProductParams{
		DiscountID: discountID,
		ProductID:  productID,
	}

	if err := s.dao.AddDiscountProduct(ctx, params); err != nil {
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

	if err := s.dao.RemoveDiscountProduct(ctx, params); err != nil {
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

	if err := s.dao.AddDiscountCategory(ctx, params); err != nil {
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

	if err := s.dao.RemoveDiscountCategory(ctx, params); err != nil {
		return fmt.Errorf("failed to remove category from discount: %w", err)
	}

	return nil
}
