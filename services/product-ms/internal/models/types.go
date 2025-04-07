package models

import "github.com/google/uuid"

// ProductID represents a product's unique identifier
type ProductID string

// NewProductID creates a new ProductID from a string
func NewProductID(id string) (ProductID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", err
	}
	return ProductID(id), nil
}

// StoreID represents a store's unique identifier
type StoreID string

// NewStoreID creates a new StoreID from a string
func NewStoreID(id string) (StoreID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", err
	}
	return StoreID(id), nil
}

// StoreStatus represents the status of a store
type StoreStatus string

const (
	StoreStatusActive   StoreStatus = "active"
	StoreStatusInactive StoreStatus = "inactive"
	StoreStatusPending  StoreStatus = "pending"
)

// ProductType represents the type of a product
type ProductType string

const (
	ProductTypeSimple   ProductType = "simple"
	ProductTypeVariable ProductType = "variable"
	ProductTypePhysical ProductType = "physical"
	ProductTypeDigital  ProductType = "digital"
	ProductTypeService  ProductType = "service"
)

// ProductStatus represents the status of a product
type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusInactive ProductStatus = "inactive"
	ProductStatusDraft    ProductStatus = "draft"
	ProductStatusArchived ProductStatus = "archived"
)

// InventoryTracking represents the inventory tracking status
type InventoryTracking string

const (
	InventoryTrackingEnabled  InventoryTracking = "enabled"
	InventoryTrackingDisabled InventoryTracking = "disabled"
)

// DiscountType represents the type of a discount
type DiscountType string

const (
	DiscountTypePercentage  DiscountType = "percentage"
	DiscountTypeFixedAmount DiscountType = "fixed_amount"
)

// DiscountScope represents the scope of a discount
type DiscountScope string

const (
	DiscountScopeAllProducts        DiscountScope = "all_products"
	DiscountScopeSpecificProducts   DiscountScope = "specific_products"
	DiscountScopeSpecificCategories DiscountScope = "specific_categories"
)

// VoucherType represents the type of voucher
type VoucherType string

const (
	VoucherTypePercentage   VoucherType = "percentage"
	VoucherTypeFixedAmount  VoucherType = "fixed_amount"
	VoucherTypeFreeShipping VoucherType = "free_shipping"
)

// VoucherStatus represents the status of a voucher
type VoucherStatus string

const (
	VoucherStatusActive   VoucherStatus = "active"
	VoucherStatusInactive VoucherStatus = "inactive"
	VoucherStatusExpired  VoucherStatus = "expired"
)
