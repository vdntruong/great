package models

import (
	"time"

	"github.com/google/uuid"
)

// Discount represents a discount in the system
type Discount struct {
	ID                uuid.UUID     `json:"id"`
	StoreID           uuid.UUID     `json:"store_id"`
	Name              string        `json:"name"`
	Code              string        `json:"code"`
	Type              DiscountType  `json:"type"`
	Value             float64       `json:"value"`
	Scope             DiscountScope `json:"scope"`
	StartDate         time.Time     `json:"start_date"`
	EndDate           *time.Time    `json:"end_date,omitempty"`
	MinPurchaseAmount *float64      `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64      `json:"max_discount_amount,omitempty"`
	UsageLimit        *int32        `json:"usage_limit,omitempty"`
	UsageCount        int32         `json:"usage_count"`
	IsActive          bool          `json:"is_active"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

// CreateDiscountParams represents the parameters for creating a discount
type CreateDiscountParams struct {
	StoreID           uuid.UUID     `json:"store_id" validate:"required"`
	Name              string        `json:"name" validate:"required,min=3,max=100"`
	Code              string        `json:"code" validate:"required,min=3,max=50"`
	Type              DiscountType  `json:"type" validate:"required,oneof=percentage fixed_amount"`
	Value             float64       `json:"value" validate:"required,min=0"`
	Scope             DiscountScope `json:"scope" validate:"required,oneof=all_products specific_products specific_categories"`
	StartDate         time.Time     `json:"start_date" validate:"required"`
	EndDate           *time.Time    `json:"end_date,omitempty"`
	MinPurchaseAmount *float64      `json:"min_purchase_amount,omitempty" validate:"omitempty,min=0"`
	MaxDiscountAmount *float64      `json:"max_discount_amount,omitempty" validate:"omitempty,min=0"`
	UsageLimit        *int32        `json:"usage_limit,omitempty" validate:"omitempty,min=1"`
	IsActive          bool          `json:"is_active"`
}

// UpdateDiscountParams represents the parameters for updating a discount
type UpdateDiscountParams struct {
	ID                uuid.UUID      `json:"id" validate:"required"`
	Name              *string        `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Code              *string        `json:"code,omitempty" validate:"omitempty,min=3,max=50"`
	Type              *DiscountType  `json:"type,omitempty" validate:"omitempty,oneof=percentage fixed_amount"`
	Value             *float64       `json:"value,omitempty" validate:"omitempty,min=0"`
	Scope             *DiscountScope `json:"scope,omitempty" validate:"omitempty,oneof=all_products specific_products specific_categories"`
	StartDate         *time.Time     `json:"start_date,omitempty"`
	EndDate           *time.Time     `json:"end_date,omitempty"`
	MinPurchaseAmount *float64       `json:"min_purchase_amount,omitempty" validate:"omitempty,min=0"`
	MaxDiscountAmount *float64       `json:"max_discount_amount,omitempty" validate:"omitempty,min=0"`
	UsageLimit        *int32         `json:"usage_limit,omitempty" validate:"omitempty,min=1"`
	IsActive          *bool          `json:"is_active,omitempty"`
}

// ListDiscountsParams represents the parameters for listing discounts
type ListDiscountsParams struct {
	StoreID uuid.UUID `json:"store_id" validate:"required"`
	Limit   int32     `json:"limit" validate:"min=1,max=100"`
	Offset  int32     `json:"offset" validate:"min=0"`
}

// DiscountList represents a list of discounts with pagination info
type DiscountList struct {
	Discounts  []Discount `json:"discounts"`
	TotalCount int64      `json:"total_count"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
}
