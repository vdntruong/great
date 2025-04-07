package handler

import (
	"errors"
	"net/http"
	"strconv"

	commonjson "commons/http/json"

	"product-ms/internal/dto"
	"product-ms/internal/models"
	"product-ms/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Discount struct {
	service *service.DiscountServiceImpl
}

func NewDiscount(service *service.DiscountServiceImpl) *Discount {
	return &Discount{
		service: service,
	}
}

// CreateDiscount handles the creation of a new discount
func (h *Discount) CreateDiscount(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDiscountRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	// Get store ID from context (set by middleware)
	storeID, ok := r.Context().Value("store_id").(uuid.UUID)
	if !ok {
		commonjson.RespondBadRequestError(w, errors.New("store ID not found"))
		return
	}

	// Convert request to model params
	params := models.CreateDiscountParams{
		StoreID:           storeID,
		Name:              req.Name,
		Code:              req.Code,
		Type:              models.DiscountType(req.Type),
		Value:             req.Value,
		Scope:             models.DiscountScope(req.Scope),
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		MinPurchaseAmount: req.MinPurchaseAmount,
		MaxDiscountAmount: req.MaxDiscountAmount,
		UsageLimit:        req.UsageLimit,
		IsActive:          req.IsActive,
	}

	// Create discount
	discount, err := h.service.CreateDiscount(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.DiscountResponse{
		ID:                discount.ID,
		StoreID:           discount.StoreID,
		Name:              discount.Name,
		Code:              discount.Code,
		Type:              string(discount.Type),
		Value:             discount.Value,
		Scope:             string(discount.Scope),
		StartDate:         discount.StartDate,
		EndDate:           discount.EndDate,
		MinPurchaseAmount: discount.MinPurchaseAmount,
		MaxDiscountAmount: discount.MaxDiscountAmount,
		UsageLimit:        discount.UsageLimit,
		UsageCount:        discount.UsageCount,
		IsActive:          discount.IsActive,
		CreatedAt:         discount.CreatedAt,
		UpdatedAt:         discount.UpdatedAt,
	}

	commonjson.RespondCreated(w, response)
}

// GetDiscount handles retrieving a discount by ID
func (h *Discount) GetDiscount(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	discount, err := h.service.GetDiscountByID(r.Context(), id.String())
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.DiscountResponse{
		ID:                discount.ID,
		StoreID:           discount.StoreID,
		Name:              discount.Name,
		Code:              discount.Code,
		Type:              string(discount.Type),
		Value:             discount.Value,
		Scope:             string(discount.Scope),
		StartDate:         discount.StartDate,
		EndDate:           discount.EndDate,
		MinPurchaseAmount: discount.MinPurchaseAmount,
		MaxDiscountAmount: discount.MaxDiscountAmount,
		UsageLimit:        discount.UsageLimit,
		UsageCount:        discount.UsageCount,
		IsActive:          discount.IsActive,
		CreatedAt:         discount.CreatedAt,
		UpdatedAt:         discount.UpdatedAt,
	}

	commonjson.RespondOK(w, response)
}

// ListDiscounts handles retrieving a list of discounts
func (h *Discount) ListDiscounts(w http.ResponseWriter, r *http.Request) {
	// Get store ID from context (set by middleware)
	storeID, ok := r.Context().Value("store_id").(uuid.UUID)
	if !ok {
		commonjson.RespondBadRequestError(w, errors.New("store ID not found"))
		return
	}

	// Get pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Convert to model params
	params := models.ListDiscountsParams{
		StoreID: storeID,
		Limit:   int32(limit),
		Offset:  int32((page - 1) * limit),
	}

	// Get discounts
	discountList, err := h.service.ListDiscounts(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.DiscountListResponse{
		Discounts:  make([]dto.DiscountResponse, len(discountList.Discounts)),
		TotalCount: discountList.TotalCount,
		Page:       discountList.Page,
		Limit:      discountList.Limit,
	}

	for i, discount := range discountList.Discounts {
		response.Discounts[i] = dto.DiscountResponse{
			ID:                discount.ID,
			StoreID:           discount.StoreID,
			Name:              discount.Name,
			Code:              discount.Code,
			Type:              string(discount.Type),
			Value:             discount.Value,
			Scope:             string(discount.Scope),
			StartDate:         discount.StartDate,
			EndDate:           discount.EndDate,
			MinPurchaseAmount: discount.MinPurchaseAmount,
			MaxDiscountAmount: discount.MaxDiscountAmount,
			UsageLimit:        discount.UsageLimit,
			UsageCount:        discount.UsageCount,
			IsActive:          discount.IsActive,
			CreatedAt:         discount.CreatedAt,
			UpdatedAt:         discount.UpdatedAt,
		}
	}

	commonjson.RespondOK(w, response)
}

// UpdateDiscount handles updating a discount
func (h *Discount) UpdateDiscount(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	var req dto.UpdateDiscountRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	// Convert request to model params
	params := models.UpdateDiscountParams{
		ID:                id,
		Name:              req.Name,
		Code:              req.Code,
		Type:              (*models.DiscountType)(req.Type),
		Value:             req.Value,
		Scope:             (*models.DiscountScope)(req.Scope),
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		MinPurchaseAmount: req.MinPurchaseAmount,
		MaxDiscountAmount: req.MaxDiscountAmount,
		UsageLimit:        req.UsageLimit,
		IsActive:          req.IsActive,
	}

	// Update discount
	discount, err := h.service.UpdateDiscount(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.DiscountResponse{
		ID:                discount.ID,
		StoreID:           discount.StoreID,
		Name:              discount.Name,
		Code:              discount.Code,
		Type:              string(discount.Type),
		Value:             discount.Value,
		Scope:             string(discount.Scope),
		StartDate:         discount.StartDate,
		EndDate:           discount.EndDate,
		MinPurchaseAmount: discount.MinPurchaseAmount,
		MaxDiscountAmount: discount.MaxDiscountAmount,
		UsageLimit:        discount.UsageLimit,
		UsageCount:        discount.UsageCount,
		IsActive:          discount.IsActive,
		CreatedAt:         discount.CreatedAt,
		UpdatedAt:         discount.UpdatedAt,
	}

	commonjson.RespondOK(w, response)
}

// DeleteDiscount handles deleting a discount
func (h *Discount) DeleteDiscount(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	if err := h.service.DeleteDiscount(r.Context(), id.String()); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// AddDiscountProduct handles adding a product to a discount
func (h *Discount) AddDiscountProduct(w http.ResponseWriter, r *http.Request) {
	discountID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	productID, err := uuid.Parse(chi.URLParam(r, "product_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid product ID"))
		return
	}

	if err := h.service.AddDiscountProduct(r.Context(), discountID, productID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// RemoveDiscountProduct handles removing a product from a discount
func (h *Discount) RemoveDiscountProduct(w http.ResponseWriter, r *http.Request) {
	discountID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	productID, err := uuid.Parse(chi.URLParam(r, "product_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid product ID"))
		return
	}

	if err := h.service.RemoveDiscountProduct(r.Context(), discountID, productID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// AddDiscountCategory handles adding a category to a discount
func (h *Discount) AddDiscountCategory(w http.ResponseWriter, r *http.Request) {
	discountID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	categoryID, err := uuid.Parse(chi.URLParam(r, "category_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid category ID"))
		return
	}

	if err := h.service.AddDiscountCategory(r.Context(), discountID, categoryID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// RemoveDiscountCategory handles removing a category from a discount
func (h *Discount) RemoveDiscountCategory(w http.ResponseWriter, r *http.Request) {
	discountID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	categoryID, err := uuid.Parse(chi.URLParam(r, "category_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid category ID"))
		return
	}

	if err := h.service.RemoveDiscountCategory(r.Context(), discountID, categoryID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}
