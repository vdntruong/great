package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateVoucherRequest represents the request body for creating a voucher
type CreateVoucherRequest struct {
	Code              string     `json:"code"`
	Type              string     `json:"type"`
	Value             *float64   `json:"value,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	StartDate         time.Time  `json:"start_date"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	Status            string     `json:"status"`
}

// UpdateVoucherRequest represents the request body for updating a voucher
type UpdateVoucherRequest struct {
	Code              *string    `json:"code,omitempty"`
	Type              *string    `json:"type,omitempty"`
	Value             *float64   `json:"value,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	StartDate         *time.Time `json:"start_date,omitempty"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	Status            *string    `json:"status,omitempty"`
}

// VoucherResponse represents the response for a voucher
type VoucherResponse struct {
	ID                uuid.UUID  `json:"id"`
	StoreID           uuid.UUID  `json:"store_id"`
	Code              string     `json:"code"`
	Type              string     `json:"type"`
	Value             *float64   `json:"value,omitempty"`
	MinPurchaseAmount *float64   `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64   `json:"max_discount_amount,omitempty"`
	StartDate         time.Time  `json:"start_date"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	UsageLimit        *int32     `json:"usage_limit,omitempty"`
	UsageCount        int32      `json:"usage_count"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// VoucherListResponse represents the response for a list of vouchers
type VoucherListResponse struct {
	Vouchers   []VoucherResponse `json:"vouchers"`
	TotalCount int64            `json:"total_count"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
}
