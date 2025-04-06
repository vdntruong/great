package models

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product in the system
type Product struct {
	ID                uuid.UUID
	StoreID           uuid.UUID
	Name              string
	Slug              string
	Description       string
	Type              string
	Status            string
	Price             float64
	CompareAtPrice    float64
	CostPrice         float64
	SKU               string
	Barcode           string
	Weight            float64
	WeightUnit        string
	IsTaxable         bool
	IsFeatured        bool
	IsGiftCard        bool
	RequiresShipping  bool
	InventoryPolicy   string
	InventoryTracking string
	SEOTitle          string
	SEODescription    string
	Metadata          map[string]interface{}
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// CreateProductParams represents the parameters for creating a product
type CreateProductParams struct {
	StoreID           uuid.UUID
	Name              string
	Slug              string
	Description       string
	Type              string
	Status            string
	Price             float64
	CompareAtPrice    float64
	CostPrice         float64
	SKU               string
	Barcode           string
	Weight            float64
	WeightUnit        string
	IsTaxable         bool
	IsFeatured        bool
	IsGiftCard        bool
	RequiresShipping  bool
	InventoryPolicy   string
	InventoryTracking string
	SEOTitle          string
	SEODescription    string
	Metadata          map[string]interface{}
}

// ListProductsParams represents the parameters for listing products
type ListProductsParams struct {
	StoreID uuid.UUID
	Limit   int32
	Offset  int32
}

// UpdateProductParams represents the parameters for updating a product
type UpdateProductParams struct {
	ID                uuid.UUID
	Name              string
	Slug              string
	Description       string
	Type              string
	Status            string
	Price             float64
	CompareAtPrice    float64
	CostPrice         float64
	SKU               string
	Barcode           string
	Weight            float64
	WeightUnit        string
	IsTaxable         bool
	IsFeatured        bool
	IsGiftCard        bool
	RequiresShipping  bool
	InventoryPolicy   string
	InventoryTracking string
	SEOTitle          string
	SEODescription    string
	Metadata          map[string]interface{}
}
