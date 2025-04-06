package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	jsonresponse "commons/http/json"
	"product-ms/internal/dto"
	"product-ms/internal/models"
	"product-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// RegisterRoutes registers all product routes
func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.Route("/stores/{store_id}/products", func(r chi.Router) {
		r.Post("/", h.CreateProduct)
		r.Get("/", h.ListProducts)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetProduct)
			r.Put("/", h.UpdateProduct)
			r.Delete("/", h.DeleteProduct)
		})
	})
}

// CreateProduct handles the creation of a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	// Convert metadata string to map
	var metadata map[string]interface{}
	if req.Metadata != "" {
		if err := json.Unmarshal([]byte(req.Metadata), &metadata); err != nil {
			jsonresponse.RespondBadRequestError(w, err)
			return
		}
	}

	// Convert inventory tracking bool to string
	inventoryTracking := "disabled"
	if req.InventoryTracking {
		inventoryTracking = "enabled"
	}

	product, err := h.productService.CreateProduct(r.Context(), models.CreateProductParams{
		StoreID:          storeID,
		Name:             req.Name,
		Slug:             req.Slug,
		Description:      req.Description,
		Type:             req.Type,
		Status:           req.Status,
		Price:            req.Price,
		CompareAtPrice:   req.CompareAtPrice,
		CostPrice:        req.CostPrice,
		SKU:              req.SKU,
		Barcode:          req.Barcode,
		Weight:           req.Weight,
		WeightUnit:       req.WeightUnit,
		IsTaxable:        req.IsTaxable,
		IsFeatured:       req.IsFeatured,
		IsGiftCard:       req.IsGiftCard,
		RequiresShipping: req.RequiresShipping,
		InventoryPolicy:  req.InventoryPolicy,
		InventoryTracking: inventoryTracking,
		SEOTitle:         req.SEOTitle,
		SEODescription:   req.SEODescription,
		Metadata:         metadata,
	})
	if err != nil {
		jsonresponse.RespondInternalServerError(w, err)
		return
	}

	// Convert metadata map to string for response
	metadataStr := ""
	if product.Metadata != nil {
		metadataBytes, err := json.Marshal(product.Metadata)
		if err != nil {
			jsonresponse.RespondInternalServerError(w, err)
			return
		}
		metadataStr = string(metadataBytes)
	}

	// Convert inventory tracking string to bool for response
	inventoryTrackingBool := product.InventoryTracking == "enabled"

	jsonresponse.RespondCreated(w, dto.ProductResponse{
		ID:               product.ID,
		StoreID:          product.StoreID,
		Name:             product.Name,
		Slug:             product.Slug,
		Description:      product.Description,
		Type:             product.Type,
		Status:           product.Status,
		Price:            product.Price,
		CompareAtPrice:   product.CompareAtPrice,
		CostPrice:        product.CostPrice,
		SKU:              product.SKU,
		Barcode:          product.Barcode,
		Weight:           product.Weight,
		WeightUnit:       product.WeightUnit,
		IsTaxable:        product.IsTaxable,
		IsFeatured:       product.IsFeatured,
		IsGiftCard:       product.IsGiftCard,
		RequiresShipping: product.RequiresShipping,
		InventoryPolicy:  product.InventoryPolicy,
		InventoryTracking: inventoryTrackingBool,
		SEOTitle:         product.SEOTitle,
		SEODescription:   product.SEODescription,
		Metadata:         metadataStr,
		CreatedAt:        product.CreatedAt,
		UpdatedAt:        product.UpdatedAt,
	})
}

// GetProduct handles retrieving a product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	product, err := h.productService.GetProduct(r.Context(), productID)
	if err != nil {
		jsonresponse.RespondInternalServerError(w, err)
		return
	}

	// Convert metadata map to string for response
	metadataStr := ""
	if product.Metadata != nil {
		metadataBytes, err := json.Marshal(product.Metadata)
		if err != nil {
			jsonresponse.RespondInternalServerError(w, err)
			return
		}
		metadataStr = string(metadataBytes)
	}

	// Convert inventory tracking string to bool for response
	inventoryTrackingBool := product.InventoryTracking == "enabled"

	jsonresponse.RespondOK(w, dto.ProductResponse{
		ID:               product.ID,
		StoreID:          product.StoreID,
		Name:             product.Name,
		Slug:             product.Slug,
		Description:      product.Description,
		Type:             product.Type,
		Status:           product.Status,
		Price:            product.Price,
		CompareAtPrice:   product.CompareAtPrice,
		CostPrice:        product.CostPrice,
		SKU:              product.SKU,
		Barcode:          product.Barcode,
		Weight:           product.Weight,
		WeightUnit:       product.WeightUnit,
		IsTaxable:        product.IsTaxable,
		IsFeatured:       product.IsFeatured,
		IsGiftCard:       product.IsGiftCard,
		RequiresShipping: product.RequiresShipping,
		InventoryPolicy:  product.InventoryPolicy,
		InventoryTracking: inventoryTrackingBool,
		SEOTitle:         product.SEOTitle,
		SEODescription:   product.SEODescription,
		Metadata:         metadataStr,
		CreatedAt:        product.CreatedAt,
		UpdatedAt:        product.UpdatedAt,
	})
}

// ListProducts handles retrieving a list of products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	products, err := h.productService.ListProducts(r.Context(), models.ListProductsParams{
		StoreID: storeID,
		Limit:   int32(limit),
		Offset:  int32(offset),
	})
	if err != nil {
		jsonresponse.RespondInternalServerError(w, err)
		return
	}

	response := dto.ListProductsResponse{
		Products: make([]dto.ProductResponse, len(products)),
		Total:    int64(len(products)),
	}

	for i, product := range products {
		// Convert metadata map to string for response
		metadataStr := ""
		if product.Metadata != nil {
			metadataBytes, err := json.Marshal(product.Metadata)
			if err != nil {
				jsonresponse.RespondInternalServerError(w, err)
				return
			}
			metadataStr = string(metadataBytes)
		}

		// Convert inventory tracking string to bool for response
		inventoryTrackingBool := product.InventoryTracking == "enabled"

		response.Products[i] = dto.ProductResponse{
			ID:               product.ID,
			StoreID:          product.StoreID,
			Name:             product.Name,
			Slug:             product.Slug,
			Description:      product.Description,
			Type:             product.Type,
			Status:           product.Status,
			Price:            product.Price,
			CompareAtPrice:   product.CompareAtPrice,
			CostPrice:        product.CostPrice,
			SKU:              product.SKU,
			Barcode:          product.Barcode,
			Weight:           product.Weight,
			WeightUnit:       product.WeightUnit,
			IsTaxable:        product.IsTaxable,
			IsFeatured:       product.IsFeatured,
			IsGiftCard:       product.IsGiftCard,
			RequiresShipping: product.RequiresShipping,
			InventoryPolicy:  product.InventoryPolicy,
			InventoryTracking: inventoryTrackingBool,
			SEOTitle:         product.SEOTitle,
			SEODescription:   product.SEODescription,
			Metadata:         metadataStr,
			CreatedAt:        product.CreatedAt,
			UpdatedAt:        product.UpdatedAt,
		}
	}

	jsonresponse.RespondOK(w, response)
}

// UpdateProduct handles updating a product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	// Convert metadata string to map
	var metadata map[string]interface{}
	if req.Metadata != nil && *req.Metadata != "" {
		if err := json.Unmarshal([]byte(*req.Metadata), &metadata); err != nil {
			jsonresponse.RespondBadRequestError(w, err)
			return
		}
	}

	// Convert inventory tracking bool to string
	var inventoryTracking string
	if req.InventoryTracking != nil {
		if *req.InventoryTracking {
			inventoryTracking = "enabled"
		} else {
			inventoryTracking = "disabled"
		}
	}

	// Convert pointer fields to non-pointer fields
	params := models.UpdateProductParams{
		ID:               productID,
		Metadata:         metadata,
		InventoryTracking: inventoryTracking,
	}

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
	if req.SEOTitle != nil {
		params.SEOTitle = *req.SEOTitle
	}
	if req.SEODescription != nil {
		params.SEODescription = *req.SEODescription
	}

	product, err := h.productService.UpdateProduct(r.Context(), params)
	if err != nil {
		jsonresponse.RespondInternalServerError(w, err)
		return
	}

	// Convert metadata map to string for response
	metadataStr := ""
	if product.Metadata != nil {
		metadataBytes, err := json.Marshal(product.Metadata)
		if err != nil {
			jsonresponse.RespondInternalServerError(w, err)
			return
		}
		metadataStr = string(metadataBytes)
	}

	// Convert inventory tracking string to bool for response
	inventoryTrackingBool := product.InventoryTracking == "enabled"

	jsonresponse.RespondOK(w, dto.ProductResponse{
		ID:               product.ID,
		StoreID:          product.StoreID,
		Name:             product.Name,
		Slug:             product.Slug,
		Description:      product.Description,
		Type:             product.Type,
		Status:           product.Status,
		Price:            product.Price,
		CompareAtPrice:   product.CompareAtPrice,
		CostPrice:        product.CostPrice,
		SKU:              product.SKU,
		Barcode:          product.Barcode,
		Weight:           product.Weight,
		WeightUnit:       product.WeightUnit,
		IsTaxable:        product.IsTaxable,
		IsFeatured:       product.IsFeatured,
		IsGiftCard:       product.IsGiftCard,
		RequiresShipping: product.RequiresShipping,
		InventoryPolicy:  product.InventoryPolicy,
		InventoryTracking: inventoryTrackingBool,
		SEOTitle:         product.SEOTitle,
		SEODescription:   product.SEODescription,
		Metadata:         metadataStr,
		CreatedAt:        product.CreatedAt,
		UpdatedAt:        product.UpdatedAt,
	})
}

// DeleteProduct handles deleting a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		jsonresponse.RespondBadRequestError(w, err)
		return
	}

	err = h.productService.DeleteProduct(r.Context(), productID)
	if err != nil {
		jsonresponse.RespondInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
