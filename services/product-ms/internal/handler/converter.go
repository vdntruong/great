package handler

import (
	"product-ms/internal/dto"
	"product-ms/internal/models"
)

// ConvertCreateProductRequestToModel converts a CreateProductRequest to a CreateProductParams
func ConvertCreateProductRequestToModel(req *dto.CreateProductRequest) models.CreateProductParams {
	return models.CreateProductParams{
		Name:              req.Name,
		Slug:              req.Slug,
		Description:       req.Description,
		Type:              req.Type,
		Status:            req.Status,
		Price:             req.Price,
		CompareAtPrice:    req.CompareAtPrice,
		CostPrice:         req.CostPrice,
		SKU:               req.SKU,
		Barcode:           req.Barcode,
		Weight:            req.Weight,
		WeightUnit:        req.WeightUnit,
		IsTaxable:         req.IsTaxable,
		IsFeatured:        req.IsFeatured,
		IsGiftCard:        req.IsGiftCard,
		RequiresShipping:  req.RequiresShipping,
		InventoryPolicy:   req.InventoryPolicy,
		InventoryTracking: req.InventoryTracking,
		SEOTitle:          req.SEOTitle,
		SEODescription:    req.SEODescription,
		Metadata:          req.Metadata,
	}
}

// ConvertUpdateProductRequestToModel converts an UpdateProductRequest to an UpdateProductParams
func ConvertUpdateProductRequestToModel(req *dto.UpdateProductRequest) models.UpdateProductParams {
	params := models.UpdateProductParams{}

	if req.Name != nil {
		params.Name = *req.Name
	}
	if req.Slug != nil {
		params.Slug = *req.Slug
	}
	if req.Description != nil {
		params.Description = *req.Description
	}
	if req.Type != nil {
		params.Type = *req.Type
	}
	if req.Status != nil {
		params.Status = *req.Status
	}
	if req.Price != nil {
		params.Price = *req.Price
	}
	if req.CompareAtPrice != nil {
		params.CompareAtPrice = *req.CompareAtPrice
	}
	if req.CostPrice != nil {
		params.CostPrice = *req.CostPrice
	}
	if req.SKU != nil {
		params.SKU = *req.SKU
	}
	if req.Barcode != nil {
		params.Barcode = *req.Barcode
	}
	if req.Weight != nil {
		params.Weight = *req.Weight
	}
	if req.WeightUnit != nil {
		params.WeightUnit = *req.WeightUnit
	}
	if req.IsTaxable != nil {
		params.IsTaxable = *req.IsTaxable
	}
	if req.IsFeatured != nil {
		params.IsFeatured = *req.IsFeatured
	}
	if req.IsGiftCard != nil {
		params.IsGiftCard = *req.IsGiftCard
	}
	if req.RequiresShipping != nil {
		params.RequiresShipping = *req.RequiresShipping
	}
	if req.InventoryPolicy != nil {
		params.InventoryPolicy = *req.InventoryPolicy
	}
	if req.InventoryTracking != nil {
		params.InventoryTracking = *req.InventoryTracking
	}
	if req.SEOTitle != nil {
		params.SEOTitle = *req.SEOTitle
	}
	if req.SEODescription != nil {
		params.SEODescription = *req.SEODescription
	}
	if req.Metadata != nil {
		params.Metadata = *req.Metadata
	}

	return params
}

// ConvertProductModelToResponse converts a Product model to a ProductResponse
func ConvertProductModelToResponse(product *models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:                product.ID,
		StoreID:           product.StoreID,
		Name:              product.Name,
		Slug:              product.Slug,
		Description:       product.Description,
		Type:              product.Type,
		Status:            product.Status,
		Price:             product.Price,
		CompareAtPrice:    product.CompareAtPrice,
		CostPrice:         product.CostPrice,
		SKU:               product.SKU,
		Barcode:           product.Barcode,
		Weight:            product.Weight,
		WeightUnit:        product.WeightUnit,
		IsTaxable:         product.IsTaxable,
		IsFeatured:        product.IsFeatured,
		IsGiftCard:        product.IsGiftCard,
		RequiresShipping:  product.RequiresShipping,
		InventoryPolicy:   product.InventoryPolicy,
		InventoryTracking: product.InventoryTracking,
		SEOTitle:          product.SEOTitle,
		SEODescription:    product.SEODescription,
		Metadata:          product.Metadata,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
	}
}

// ConvertCreateStoreRequestToModel converts a CreateStoreRequest to a CreateStoreParams
func ConvertCreateStoreRequestToModel(req *dto.CreateStoreRequest) models.CreateStoreParams {
	return models.CreateStoreParams{
		Name:         req.Name,
		Slug:         req.Slug,
		Description:  req.Description,
		LogoURL:      req.LogoURL,
		CoverURL:     req.CoverURL,
		Status:       req.Status,
		IsVerified:   req.IsVerified,
		OwnerID:      req.OwnerID,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Address:      req.Address,
		Settings:     req.Settings,
	}
}

// ConvertUpdateStoreRequestToModel converts an UpdateStoreRequest to an UpdateStoreParams
func ConvertUpdateStoreRequestToModel(req *dto.UpdateStoreRequest) models.UpdateStoreParams {
	params := models.UpdateStoreParams{}

	if req.Name != nil {
		params.Name = *req.Name
	}
	if req.Slug != nil {
		params.Slug = *req.Slug
	}
	if req.Description != nil {
		params.Description = *req.Description
	}
	if req.LogoURL != nil {
		params.LogoURL = *req.LogoURL
	}
	if req.CoverURL != nil {
		params.CoverURL = *req.CoverURL
	}
	if req.Status != nil {
		params.Status = *req.Status
	}
	if req.IsVerified != nil {
		params.IsVerified = *req.IsVerified
	}
	if req.OwnerID != nil {
		params.OwnerID = *req.OwnerID
	}
	if req.ContactEmail != nil {
		params.ContactEmail = *req.ContactEmail
	}
	if req.ContactPhone != nil {
		params.ContactPhone = *req.ContactPhone
	}
	if req.Address != nil {
		params.Address = *req.Address
	}
	if req.Settings != nil {
		params.Settings = *req.Settings
	}

	return params
}

// ConvertStoreModelToResponse converts a Store model to a StoreResponse
func ConvertStoreModelToResponse(store *models.Store) dto.StoreResponse {
	return dto.StoreResponse{
		ID:           store.ID,
		Name:         store.Name,
		Slug:         store.Slug,
		Description:  store.Description,
		LogoURL:      store.LogoURL,
		CoverURL:     store.CoverURL,
		Status:       store.Status,
		IsVerified:   store.IsVerified,
		OwnerID:      store.OwnerID,
		ContactEmail: store.ContactEmail,
		ContactPhone: store.ContactPhone,
		Address:      store.Address,
		Settings:     store.Settings,
		CreatedAt:    store.CreatedAt,
		UpdatedAt:    store.UpdatedAt,
	}
}
