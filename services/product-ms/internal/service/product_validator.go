package service

import (
	"fmt"
	"product-ms/internal/models"
	"strings"
)

// ValidateCreateProductParams validates the parameters for creating a product
func ValidateCreateProductParams(params models.CreateProductParams) error {
	if params.StoreID.String() == "" {
		return fmt.Errorf("store_id is required")
	}

	if strings.TrimSpace(params.Name) == "" {
		return fmt.Errorf("name is required")
	}

	if strings.TrimSpace(params.Slug) == "" {
		return fmt.Errorf("slug is required")
	}

	if params.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	if params.CompareAtPrice < 0 {
		return fmt.Errorf("compare_at_price cannot be negative")
	}

	if params.CostPrice < 0 {
		return fmt.Errorf("cost_price cannot be negative")
	}

	if params.Weight < 0 {
		return fmt.Errorf("weight cannot be negative")
	}

	// Validate product type
	if !isValidProductType(params.Type) {
		return fmt.Errorf("invalid product type: %s", params.Type)
	}

	// Validate product status
	if !isValidProductStatus(params.Status) {
		return fmt.Errorf("invalid product status: %s", params.Status)
	}

	// Validate inventory tracking
	if params.InventoryTracking != "" && params.InventoryTracking != "enabled" && params.InventoryTracking != "disabled" {
		return fmt.Errorf("invalid inventory tracking value: %s", params.InventoryTracking)
	}

	return nil
}

// ValidateUpdateProductParams validates the parameters for updating a product
func ValidateUpdateProductParams(params models.UpdateProductParams) error {
	if params.ID.String() == "" {
		return fmt.Errorf("id is required")
	}

	if strings.TrimSpace(params.Name) == "" {
		return fmt.Errorf("name is required")
	}

	if strings.TrimSpace(params.Slug) == "" {
		return fmt.Errorf("slug is required")
	}

	if params.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}

	if params.CompareAtPrice < 0 {
		return fmt.Errorf("compare_at_price cannot be negative")
	}

	if params.CostPrice < 0 {
		return fmt.Errorf("cost_price cannot be negative")
	}

	if params.Weight < 0 {
		return fmt.Errorf("weight cannot be negative")
	}

	// Validate product type
	if !isValidProductType(params.Type) {
		return fmt.Errorf("invalid product type: %s", params.Type)
	}

	// Validate product status
	if !isValidProductStatus(params.Status) {
		return fmt.Errorf("invalid product status: %s", params.Status)
	}

	// Validate inventory tracking
	if params.InventoryTracking != "" && params.InventoryTracking != "enabled" && params.InventoryTracking != "disabled" {
		return fmt.Errorf("invalid inventory tracking value: %s", params.InventoryTracking)
	}

	return nil
}

// isValidProductType checks if the product type is valid
func isValidProductType(productType string) bool {
	switch productType {
	case "physical", "digital", "service", "gift_card":
		return true
	default:
		return false
	}
}

// isValidProductStatus checks if the product status is valid
func isValidProductStatus(status string) bool {
	switch status {
	case "draft", "published", "archived":
		return true
	default:
		return false
	}
}
