package models

import (
	"time"

	"github.com/google/uuid"
)

// Voucher represents a voucher in the system
type Voucher struct {
	ID                uuid.UUID     `json:"id"`
	StoreID           uuid.UUID     `json:"store_id"`
	Code              string        `json:"code"`
	Type              VoucherType   `json:"type"`
	Value             *float64      `json:"value,omitempty"`
	MinPurchaseAmount *float64      `json:"min_purchase_amount,omitempty"`
	MaxDiscountAmount *float64      `json:"max_discount_amount,omitempty"`
	StartDate         time.Time     `json:"start_date"`
	EndDate           *time.Time    `json:"end_date,omitempty"`
	UsageLimit        *int32        `json:"usage_limit,omitempty"`
	UsageCount        int32         `json:"usage_count"`
	Status            VoucherStatus `json:"status"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

// CreateVoucherParams represents the parameters for creating a voucher
type CreateVoucherParams struct {
	StoreID           uuid.UUID     `json:"store_id" validate:"required"`
	Code              string        `json:"code" validate:"required,min=3,max=50"`
	Type              VoucherType   `json:"type" validate:"required,oneof=percentage fixed_amount free_shipping"`
	Value             *float64      `json:"value,omitempty" validate:"required_if=Type percentage,fixed_amount,omitempty,min=0"`
	MinPurchaseAmount *float64      `json:"min_purchase_amount,omitempty" validate:"omitempty,min=0"`
	MaxDiscountAmount *float64      `json:"max_discount_amount,omitempty" validate:"omitempty,min=0"`
	StartDate         time.Time     `json:"start_date" validate:"required"`
	EndDate           *time.Time    `json:"end_date,omitempty"`
	UsageLimit        *int32        `json:"usage_limit,omitempty" validate:"omitempty,min=1"`
	Status            VoucherStatus `json:"status" validate:"required,oneof=active inactive"`
}

// UpdateVoucherParams represents the parameters for updating a voucher
type UpdateVoucherParams struct {
	ID                uuid.UUID      `json:"id" validate:"required"`
	Code              *string        `json:"code,omitempty" validate:"omitempty,min=3,max=50"`
	Type              *VoucherType   `json:"type,omitempty" validate:"omitempty,oneof=percentage fixed_amount free_shipping"`
	Value             *float64       `json:"value,omitempty" validate:"omitempty,min=0"`
	MinPurchaseAmount *float64       `json:"min_purchase_amount,omitempty" validate:"omitempty,min=0"`
	MaxDiscountAmount *float64       `json:"max_discount_amount,omitempty" validate:"omitempty,min=0"`
	StartDate         *time.Time     `json:"start_date,omitempty"`
	EndDate           *time.Time     `json:"end_date,omitempty"`
	UsageLimit        *int32         `json:"usage_limit,omitempty" validate:"omitempty,min=1"`
	Status            *VoucherStatus `json:"status,omitempty" validate:"omitempty,oneof=active inactive expired"`
}

// ListVouchersParams represents the parameters for listing vouchers
type ListVouchersParams struct {
	StoreID uuid.UUID `json:"store_id" validate:"required"`
	Limit   int32     `json:"limit" validate:"min=1,max=100"`
	Offset  int32     `json:"offset" validate:"min=0"`
}

// VoucherList represents a list of vouchers with pagination info
type VoucherList struct {
	Vouchers   []Voucher `json:"vouchers"`
	TotalCount int64     `json:"total_count"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}
