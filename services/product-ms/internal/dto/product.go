package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
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
}

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name              *string                 `json:"name,omitempty"`
	Slug              *string                 `json:"slug,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	Price             *float64                `json:"price,omitempty"`
	CompareAtPrice    *float64                `json:"compare_at_price,omitempty"`
	CostPrice         *float64                `json:"cost_price,omitempty"`
	SKU               *string                 `json:"sku,omitempty"`
	Barcode           *string                 `json:"barcode,omitempty"`
	Weight            *float64                `json:"weight,omitempty"`
	WeightUnit        *string                 `json:"weight_unit,omitempty"`
	IsTaxable         *bool                   `json:"is_taxable,omitempty"`
	IsFeatured        *bool                   `json:"is_featured,omitempty"`
	IsGiftCard        *bool                   `json:"is_gift_card,omitempty"`
	RequiresShipping  *bool                   `json:"requires_shipping,omitempty"`
	InventoryPolicy   *string                 `json:"inventory_policy,omitempty"`
	InventoryTracking *string                 `json:"inventory_tracking,omitempty"`
	SEOTitle          *string                 `json:"seo_title,omitempty"`
	SEODescription    *string                 `json:"seo_description,omitempty"`
	Metadata          *map[string]interface{} `json:"metadata,omitempty"`
}

// ProductResponse represents the response for a product
type ProductResponse struct {
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

// ListProductsResponse represents the response for listing products
type ListProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
}
