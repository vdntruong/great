package models

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product in the system
type Product struct {
	ID                uuid.UUID              `json:"id"`
	StoreID           uuid.UUID              `json:"store_id"`
	Name              string                 `json:"name"`
	Slug              string                 `json:"slug"`
	Description       string                 `json:"description"`
	Type              string                 `json:"type"`
	Status            string                 `json:"status"`
	Price             float64                `json:"price"`
	CompareAtPrice    float64                `json:"compare_at_price"`
	CostPrice         float64                `json:"cost_price"`
	SKU               string                 `json:"sku"`
	Barcode           string                 `json:"barcode"`
	Weight            float64                `json:"weight"`
	WeightUnit        string                 `json:"weight_unit"`
	IsTaxable         bool                   `json:"is_taxable"`
	IsFeatured        bool                   `json:"is_featured"`
	IsGiftCard        bool                   `json:"is_gift_card"`
	RequiresShipping  bool                   `json:"requires_shipping"`
	InventoryPolicy   string                 `json:"inventory_policy"`
	InventoryTracking string                 `json:"inventory_tracking"`
	SEOTitle          string                 `json:"seo_title"`
	SEODescription    string                 `json:"seo_description"`
	Metadata          map[string]interface{} `json:"metadata"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

// CreateProductParams represents the parameters for creating a product
type CreateProductParams struct {
	StoreID           uuid.UUID              `json:"store_id" validate:"required"`
	Name              string                 `json:"name" validate:"required,min=3,max=100"`
	Slug              string                 `json:"slug" validate:"required,min=3,max=100,slug"`
	Description       string                 `json:"description" validate:"max=500"`
	Type              string                 `json:"type" validate:"required,oneof=physical digital service gift_card"`
	Status            string                 `json:"status" validate:"required,oneof=draft published archived"`
	Price             float64                `json:"price" validate:"required,min=0"`
	CompareAtPrice    float64                `json:"compare_at_price" validate:"min=0"`
	CostPrice         float64                `json:"cost_price" validate:"min=0"`
	SKU               string                 `json:"sku" validate:"max=50"`
	Barcode           string                 `json:"barcode" validate:"max=50"`
	Weight            float64                `json:"weight" validate:"min=0"`
	WeightUnit        string                 `json:"weight_unit" validate:"max=10"`
	IsTaxable         bool                   `json:"is_taxable"`
	IsFeatured        bool                   `json:"is_featured"`
	IsGiftCard        bool                   `json:"is_gift_card"`
	RequiresShipping  bool                   `json:"requires_shipping"`
	InventoryPolicy   string                 `json:"inventory_policy" validate:"max=50"`
	InventoryTracking string                 `json:"inventory_tracking" validate:"oneof=enabled disabled"`
	SEOTitle          string                 `json:"seo_title" validate:"max=100"`
	SEODescription    string                 `json:"seo_description" validate:"max=200"`
	Metadata          map[string]interface{} `json:"metadata"`
}

// ListProductsParams represents the parameters for listing products
type ListProductsParams struct {
	StoreID uuid.UUID `json:"store_id" validate:"required"`
	Limit   int32     `json:"limit" validate:"min=1,max=100"`
	Offset  int32     `json:"offset" validate:"min=0"`
}

// UpdateProductParams represents the parameters for updating a product
type UpdateProductParams struct {
	ID                uuid.UUID              `json:"id" validate:"required"`
	Name              *string                `json:"name,omitempty" validate:"omitempty,min=3,max=100"`
	Slug              *string                `json:"slug,omitempty" validate:"omitempty,min=3,max=100,slug"`
	Description       *string                `json:"description,omitempty" validate:"omitempty,max=500"`
	Type              *string                `json:"type,omitempty" validate:"omitempty,oneof=physical digital service gift_card"`
	Status            *string                `json:"status,omitempty" validate:"omitempty,oneof=draft published archived"`
	Price             *float64               `json:"price,omitempty" validate:"omitempty,min=0"`
	CompareAtPrice    *float64               `json:"compare_at_price,omitempty" validate:"omitempty,min=0"`
	CostPrice         *float64               `json:"cost_price,omitempty" validate:"omitempty,min=0"`
	SKU               *string                `json:"sku,omitempty" validate:"omitempty,max=50"`
	Barcode           *string                `json:"barcode,omitempty" validate:"omitempty,max=50"`
	Weight            *float64               `json:"weight,omitempty" validate:"omitempty,min=0"`
	WeightUnit        *string                `json:"weight_unit,omitempty" validate:"omitempty,max=10"`
	IsTaxable         *bool                  `json:"is_taxable,omitempty"`
	IsFeatured        *bool                  `json:"is_featured,omitempty"`
	IsGiftCard        *bool                  `json:"is_gift_card,omitempty"`
	RequiresShipping  *bool                  `json:"requires_shipping,omitempty"`
	InventoryPolicy   *string                `json:"inventory_policy,omitempty" validate:"omitempty,max=50"`
	InventoryTracking *string                `json:"inventory_tracking,omitempty" validate:"omitempty,oneof=enabled disabled"`
	SEOTitle          *string                `json:"seo_title,omitempty" validate:"omitempty,max=100"`
	SEODescription    *string                `json:"seo_description,omitempty" validate:"omitempty,max=200"`
	Metadata          *map[string]interface{} `json:"metadata,omitempty"`
}
