package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"product-ms/internal/models"
	"product-ms/internal/repository/dao"

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

// ConvertCreateProductParamsToDAO converts a model CreateProductParams to DAO CreateProductParams
func ConvertCreateProductParamsToDAO(params models.CreateProductParams) (*dao.CreateProductParams, error) {
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

// ConvertUpdateProductParamsToDAO converts a model UpdateProductParams to DAO UpdateProductParams
func ConvertUpdateProductParamsToDAO(params models.UpdateProductParams) (*dao.UpdateProductParams, error) {
	var metadata pqtype.NullRawMessage
	if params.Metadata != nil {
		metadataBytes, err := json.Marshal(*params.Metadata)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		metadata = pqtype.NullRawMessage{
			RawMessage: metadataBytes,
			Valid:      true,
		}
	}

	daoParams := &dao.UpdateProductParams{
		ID:                params.ID,
		InventoryQuantity: sql.NullInt32{Int32: 0, Valid: false},
	}

	if params.Name != nil {
		daoParams.Name = *params.Name
	}
	if params.Slug != nil {
		daoParams.Slug = *params.Slug
	}
	if params.Description != nil {
		daoParams.Description = sql.NullString{String: *params.Description, Valid: true}
	}
	if params.Type != nil {
		daoParams.Type = dao.ProductType(*params.Type)
	}
	if params.Status != nil {
		daoParams.Status = dao.ProductStatus(*params.Status)
	}
	if params.Price != nil {
		daoParams.Price = fmt.Sprintf("%.2f", *params.Price)
	}
	if params.CompareAtPrice != nil {
		daoParams.CompareAtPrice = sql.NullString{String: fmt.Sprintf("%.2f", *params.CompareAtPrice), Valid: true}
	}
	if params.CostPrice != nil {
		daoParams.CostPrice = sql.NullString{String: fmt.Sprintf("%.2f", *params.CostPrice), Valid: true}
	}
	if params.SKU != nil {
		daoParams.Sku = sql.NullString{String: *params.SKU, Valid: true}
	}
	if params.Barcode != nil {
		daoParams.Barcode = sql.NullString{String: *params.Barcode, Valid: true}
	}
	if params.Weight != nil {
		daoParams.Weight = sql.NullString{String: fmt.Sprintf("%.2f", *params.Weight), Valid: true}
	}
	if params.WeightUnit != nil {
		daoParams.WeightUnit = sql.NullString{String: *params.WeightUnit, Valid: true}
	}
	if params.IsTaxable != nil {
		daoParams.IsTaxable = sql.NullBool{Bool: *params.IsTaxable, Valid: true}
	}
	if params.IsFeatured != nil {
		daoParams.IsFeatured = sql.NullBool{Bool: *params.IsFeatured, Valid: true}
	}
	if params.IsGiftCard != nil {
		daoParams.IsGiftCard = sql.NullBool{Bool: *params.IsGiftCard, Valid: true}
	}
	if params.RequiresShipping != nil {
		daoParams.RequiresShipping = sql.NullBool{Bool: *params.RequiresShipping, Valid: true}
	}
	if params.InventoryPolicy != nil {
		daoParams.InventoryPolicy = sql.NullString{String: *params.InventoryPolicy, Valid: true}
	}
	if params.InventoryTracking != nil {
		daoParams.InventoryTracking = sql.NullBool{Bool: *params.InventoryTracking == "enabled", Valid: true}
	}
	if params.SEOTitle != nil {
		daoParams.SeoTitle = sql.NullString{String: *params.SEOTitle, Valid: true}
	}
	if params.SEODescription != nil {
		daoParams.SeoDescription = sql.NullString{String: *params.SEODescription, Valid: true}
	}
	daoParams.Metadata = metadata

	return daoParams, nil
}

// ConvertProductToModel converts a DAO product to a model product
func ConvertProductToModel(product *dao.Product) *models.Product {
	if product == nil {
		return nil
	}

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
		ID:                product.ID,
		StoreID:           product.StoreID,
		Name:              product.Name,
		Slug:              product.Slug,
		Description:       product.Description.String,
		Type:              string(product.Type),
		Status:            string(product.Status),
		Price:             price,
		CompareAtPrice:    compareAtPrice,
		CostPrice:         costPrice,
		SKU:               product.Sku.String,
		Barcode:           product.Barcode.String,
		Weight:            weight,
		WeightUnit:        product.WeightUnit.String,
		IsTaxable:         product.IsTaxable.Bool,
		IsFeatured:        product.IsFeatured.Bool,
		IsGiftCard:        product.IsGiftCard.Bool,
		RequiresShipping:  product.RequiresShipping.Bool,
		InventoryPolicy:   product.InventoryPolicy.String,
		InventoryTracking: inventoryTracking,
		SEOTitle:          product.SeoTitle.String,
		SEODescription:    product.SeoDescription.String,
		Metadata:          metadata,
		CreatedAt:         product.CreatedAt.Time,
		UpdatedAt:         product.UpdatedAt.Time,
	}
}

// ConvertListProductsParamsToDAO converts a model ListProductsParams to DAO ListProductsParams
func ConvertListProductsParamsToDAO(params models.ListProductsParams) *dao.ListProductsParams {
	return &dao.ListProductsParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}
}
