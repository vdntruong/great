package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateDiscountRequest represents the request body for creating a discount
type CreateDiscountRequest struct {
	Name              string     `json:"name"`
	Code              string     `json:"code"`
	Type              string     `json:"type"`
	Value             float64    `json:"value"`
	Scope             string     `json:"scope"`
	StartDate         time.Time  `json:"start_date"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	IsActive          bool       `json:"is_active"`
}

// UpdateDiscountRequest represents the request body for updating a discount
type UpdateDiscountRequest struct {
	Name              *string    `json:"name,omitempty"`
	Code              *string    `json:"code,omitempty"`
	Type              *string    `json:"type,omitempty"`
	Value             *float64   `json:"value,omitempty"`
	Scope             *string    `json:"scope,omitempty"`
	StartDate         *time.Time `json:"start_date,omitempty"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	IsActive          *bool      `json:"is_active,omitempty"`
}

// DiscountResponse represents the response for a discount
type DiscountResponse struct {
	ID                uuid.UUID  `json:"id"`
	StoreID           uuid.UUID  `json:"store_id"`
	Name              string     `json:"name"`
	Code              string     `json:"code"`
	Type              string     `json:"type"`
	Value             float64    `json:"value"`
	Scope             string     `json:"scope"`
	StartDate         time.Time  `json:"start_date"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	UsageCount        int32      `json:"usage_count"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// DiscountListResponse represents the response for a list of discounts
type DiscountListResponse struct {
	Discounts   []DiscountResponse `json:"discounts"`
	TotalCount  int64             `json:"total_count"`
	Page        int               `json:"page"`
	Limit       int               `json:"limit"`
}
