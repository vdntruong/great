package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"product-ms/internal/models"
	"product-ms/internal/repository/dao"
	"strconv"

	"github.com/sqlc-dev/pqtype"
)

// convertMetadata converts map to pqtype.NullRawMessage
func convertMetadata(metadata map[string]interface{}) (pqtype.NullRawMessage, error) {
	if metadata == nil {
		return pqtype.NullRawMessage{}, nil
	}

	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return pqtype.NullRawMessage{}, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return pqtype.NullRawMessage{
		RawMessage: metadataBytes,
		Valid:      true,
	}, nil
}

// convertCreateParams converts CreateProductParams to DAO params
func convertCreateParams(params models.CreateProductParams) (*dao.CreateProductParams, error) {
	metadata, err := convertMetadata(params.Metadata)
	if err != nil {
		return nil, err
	}

	return &dao.CreateProductParams{
		StoreID:           params.StoreID,
		Name:              params.Name,
		Slug:              params.Slug,
		Description:       sql.NullString{String: params.Description, Valid: params.Description != ""},
		Type:              dao.ProductType(params.Type),
		Status:            dao.ProductStatus(params.Status),
		Price:             fmt.Sprintf("%.2f", params.Price),
		CompareAtPrice:    sql.NullString{String: fmt.Sprintf("%.2f", params.CompareAtPrice), Valid: params.CompareAtPrice != 0},
		CostPrice:         sql.NullString{String: fmt.Sprintf("%.2f", params.CostPrice), Valid: params.CostPrice != 0},
		Sku:               sql.NullString{String: params.SKU, Valid: params.SKU != ""},
		Barcode:           sql.NullString{String: params.Barcode, Valid: params.Barcode != ""},
		Weight:            sql.NullString{String: fmt.Sprintf("%.2f", params.Weight), Valid: params.Weight != 0},
		WeightUnit:        sql.NullString{String: params.WeightUnit, Valid: params.WeightUnit != ""},
		IsTaxable:         sql.NullBool{Bool: params.IsTaxable, Valid: true},
		IsFeatured:        sql.NullBool{Bool: params.IsFeatured, Valid: true},
		IsGiftCard:        sql.NullBool{Bool: params.IsGiftCard, Valid: true},
		RequiresShipping:  sql.NullBool{Bool: params.RequiresShipping, Valid: true},
		InventoryQuantity: sql.NullInt32{Int32: 0, Valid: false},
		InventoryPolicy:   sql.NullString{String: params.InventoryPolicy, Valid: params.InventoryPolicy != ""},
		InventoryTracking: sql.NullBool{Bool: params.InventoryTracking == "enabled", Valid: true},
		SeoTitle:          sql.NullString{String: params.SEOTitle, Valid: params.SEOTitle != ""},
		SeoDescription:    sql.NullString{String: params.SEODescription, Valid: params.SEODescription != ""},
		Metadata:          metadata,
	}, nil
}

// convertUpdateParams converts UpdateProductParams to DAO params
func convertUpdateParams(params models.UpdateProductParams) (*dao.UpdateProductParams, error) {
	metadata, err := convertMetadata(params.Metadata)
	if err != nil {
		return nil, err
	}

	return &dao.UpdateProductParams{
		ID:                params.ID,
		Name:              params.Name,
		Slug:              params.Slug,
		Description:       sql.NullString{String: params.Description, Valid: params.Description != ""},
		Type:              dao.ProductType(params.Type),
		Status:            dao.ProductStatus(params.Status),
		Price:             fmt.Sprintf("%.2f", params.Price),
		CompareAtPrice:    sql.NullString{String: fmt.Sprintf("%.2f", params.CompareAtPrice), Valid: params.CompareAtPrice != 0},
		CostPrice:         sql.NullString{String: fmt.Sprintf("%.2f", params.CostPrice), Valid: params.CostPrice != 0},
		Sku:               sql.NullString{String: params.SKU, Valid: params.SKU != ""},
		Barcode:           sql.NullString{String: params.Barcode, Valid: params.Barcode != ""},
		Weight:            sql.NullString{String: fmt.Sprintf("%.2f", params.Weight), Valid: params.Weight != 0},
		WeightUnit:        sql.NullString{String: params.WeightUnit, Valid: params.WeightUnit != ""},
		IsTaxable:         sql.NullBool{Bool: params.IsTaxable, Valid: true},
		IsFeatured:        sql.NullBool{Bool: params.IsFeatured, Valid: true},
		IsGiftCard:        sql.NullBool{Bool: params.IsGiftCard, Valid: true},
		RequiresShipping:  sql.NullBool{Bool: params.RequiresShipping, Valid: true},
		InventoryQuantity: sql.NullInt32{Int32: 0, Valid: false},
		InventoryPolicy:   sql.NullString{String: params.InventoryPolicy, Valid: params.InventoryPolicy != ""},
		InventoryTracking: sql.NullBool{Bool: params.InventoryTracking == "enabled", Valid: true},
		SeoTitle:          sql.NullString{String: params.SEOTitle, Valid: params.SEOTitle != ""},
		SeoDescription:    sql.NullString{String: params.SEODescription, Valid: params.SEODescription != ""},
		Metadata:          metadata,
	}, nil
}

// ConvertDAOProductToModel converts DAO product to model product
func ConvertDAOProductToModel(product *dao.Product) *models.Product {
	// Convert metadata to map
	var metadata map[string]interface{}
	if product.Metadata.Valid {
		if err := json.Unmarshal(product.Metadata.RawMessage, &metadata); err != nil {
			metadata = nil
		}
	}

	// Convert string prices to float64
	price, _ := strconv.ParseFloat(product.Price, 64)
	var compareAtPrice float64
	if product.CompareAtPrice.Valid {
		compareAtPrice, _ = strconv.ParseFloat(product.CompareAtPrice.String, 64)
	}
	var costPrice float64
	if product.CostPrice.Valid {
		costPrice, _ = strconv.ParseFloat(product.CostPrice.String, 64)
	}

	// Convert string weight to float64
	var weight float64
	if product.Weight.Valid {
		weight, _ = strconv.ParseFloat(product.Weight.String, 64)
	}

	// Convert inventory tracking to string
	inventoryTracking := "disabled"
	if product.InventoryTracking.Valid && product.InventoryTracking.Bool {
		inventoryTracking = "enabled"
	}

	return &models.Product{
		ID:               product.ID,
		StoreID:          product.StoreID,
		Name:             product.Name,
		Slug:             product.Slug,
		Description:      product.Description.String,
		Type:             string(product.Type),
		Status:           string(product.Status),
		Price:            price,
		CompareAtPrice:   compareAtPrice,
		CostPrice:        costPrice,
		SKU:              product.Sku.String,
		Barcode:          product.Barcode.String,
		Weight:           weight,
		WeightUnit:       product.WeightUnit.String,
		IsTaxable:        product.IsTaxable.Bool,
		IsFeatured:       product.IsFeatured.Bool,
		IsGiftCard:       product.IsGiftCard.Bool,
		RequiresShipping: product.RequiresShipping.Bool,
		InventoryPolicy:  product.InventoryPolicy.String,
		InventoryTracking: inventoryTracking,
		SEOTitle:         product.SeoTitle.String,
		SEODescription:   product.SeoDescription.String,
		Metadata:         metadata,
		CreatedAt:        product.CreatedAt.Time,
		UpdatedAt:        product.UpdatedAt.Time,
	}
}
