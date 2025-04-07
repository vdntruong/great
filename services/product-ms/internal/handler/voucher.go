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

type Voucher struct {
	service service.VoucherService
}

func NewVoucher(service service.VoucherService) *Voucher {
	return &Voucher{
		service: service,
	}
}

// CreateVoucher handles the creation of a new voucher
func (h *Voucher) CreateVoucher(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateVoucherRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("store ID not found"))
		return
	}

	// Convert request to model params
	params := models.CreateVoucherParams{
		StoreID:           storeID,
		Code:              req.Code,
		Type:              models.VoucherType(req.Type),
		Value:             req.Value,
		MinPurchaseAmount: req.MinPurchaseAmount,
		MaxDiscountAmount: req.MaxDiscountAmount,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		UsageLimit:        req.UsageLimit,
		Status:            models.VoucherStatus(req.Status),
	}

	// Create voucher
	voucher, err := h.service.CreateVoucher(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.VoucherResponse{
		ID:                voucher.ID,
		StoreID:           voucher.StoreID,
		Code:              voucher.Code,
		Type:              string(voucher.Type),
		Value:             voucher.Value,
		MinPurchaseAmount: voucher.MinPurchaseAmount,
		MaxDiscountAmount: voucher.MaxDiscountAmount,
		StartDate:         voucher.StartDate,
		EndDate:           voucher.EndDate,
		UsageLimit:        voucher.UsageLimit,
		UsageCount:        voucher.UsageCount,
		Status:            string(voucher.Status),
		CreatedAt:         voucher.CreatedAt,
		UpdatedAt:         voucher.UpdatedAt,
	}

	commonjson.RespondCreated(w, response)
}

// GetVoucher handles retrieving a voucher by ID
func (h *Voucher) GetVoucher(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid voucher ID"))
		return
	}

	voucher, err := h.service.GetVoucherByID(r.Context(), id.String())
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.VoucherResponse{
		ID:                voucher.ID,
		StoreID:           voucher.StoreID,
		Code:              voucher.Code,
		Type:              string(voucher.Type),
		Value:             voucher.Value,
		MinPurchaseAmount: voucher.MinPurchaseAmount,
		MaxDiscountAmount: voucher.MaxDiscountAmount,
		StartDate:         voucher.StartDate,
		EndDate:           voucher.EndDate,
		UsageLimit:        voucher.UsageLimit,
		UsageCount:        voucher.UsageCount,
		Status:            string(voucher.Status),
		CreatedAt:         voucher.CreatedAt,
		UpdatedAt:         voucher.UpdatedAt,
	}

	commonjson.RespondOK(w, response)
}

// ListVouchers handles retrieving a list of vouchers
func (h *Voucher) ListVouchers(w http.ResponseWriter, r *http.Request) {
	storeID, err := uuid.Parse(chi.URLParam(r, "store_id"))
	if err != nil {
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
	params := models.ListVouchersParams{
		StoreID: storeID,
		Limit:   int32(limit),
		Offset:  int32((page - 1) * limit),
	}

	// Get vouchers
	voucherList, err := h.service.ListVouchers(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.VoucherListResponse{
		Vouchers:   make([]dto.VoucherResponse, len(voucherList.Vouchers)),
		TotalCount: voucherList.TotalCount,
		Page:       voucherList.Page,
		Limit:      voucherList.Limit,
	}

	for i, voucher := range voucherList.Vouchers {
		response.Vouchers[i] = dto.VoucherResponse{
			ID:                voucher.ID,
			StoreID:           voucher.StoreID,
			Code:              voucher.Code,
			Type:              string(voucher.Type),
			Value:             voucher.Value,
			MinPurchaseAmount: voucher.MinPurchaseAmount,
			MaxDiscountAmount: voucher.MaxDiscountAmount,
			StartDate:         voucher.StartDate,
			EndDate:           voucher.EndDate,
			UsageLimit:        voucher.UsageLimit,
			UsageCount:        voucher.UsageCount,
			Status:            string(voucher.Status),
			CreatedAt:         voucher.CreatedAt,
			UpdatedAt:         voucher.UpdatedAt,
		}
	}

	commonjson.RespondOK(w, response)
}

// UpdateVoucher handles updating a voucher
func (h *Voucher) UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid voucher ID"))
		return
	}

	var req dto.UpdateVoucherRequest
	if err := commonjson.DecodeRequest(r, &req); err != nil {
		commonjson.RespondBadRequestError(w, err)
		return
	}

	// Convert request to model params
	params := models.UpdateVoucherParams{
		ID:                id,
		Code:              req.Code,
		Type:              (*models.VoucherType)(req.Type),
		Value:             req.Value,
		MinPurchaseAmount: req.MinPurchaseAmount,
		MaxDiscountAmount: req.MaxDiscountAmount,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		UsageLimit:        req.UsageLimit,
		Status:            (*models.VoucherStatus)(req.Status),
	}

	// Update voucher
	voucher, err := h.service.UpdateVoucher(r.Context(), params)
	if err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	// Convert to response
	response := dto.VoucherResponse{
		ID:                voucher.ID,
		StoreID:           voucher.StoreID,
		Code:              voucher.Code,
		Type:              string(voucher.Type),
		Value:             voucher.Value,
		MinPurchaseAmount: voucher.MinPurchaseAmount,
		MaxDiscountAmount: voucher.MaxDiscountAmount,
		StartDate:         voucher.StartDate,
		EndDate:           voucher.EndDate,
		UsageLimit:        voucher.UsageLimit,
		UsageCount:        voucher.UsageCount,
		Status:            string(voucher.Status),
		CreatedAt:         voucher.CreatedAt,
		UpdatedAt:         voucher.UpdatedAt,
	}

	commonjson.RespondOK(w, response)
}

// DeleteVoucher handles deleting a voucher
func (h *Voucher) DeleteVoucher(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid voucher ID"))
		return
	}

	if err := h.service.DeleteVoucher(r.Context(), id.String()); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// AddVoucherProduct handles adding a product to a voucher
func (h *Voucher) AddVoucherProduct(w http.ResponseWriter, r *http.Request) {
	voucherID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	productID, err := uuid.Parse(chi.URLParam(r, "product_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid product ID"))
		return
	}

	if err := h.service.AddVoucherProduct(r.Context(), voucherID, productID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// RemoveVoucherProduct handles removing a product from a voucher
func (h *Voucher) RemoveVoucherProduct(w http.ResponseWriter, r *http.Request) {
	voucherID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	productID, err := uuid.Parse(chi.URLParam(r, "product_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid product ID"))
		return
	}

	if err := h.service.RemoveVoucherProduct(r.Context(), voucherID, productID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// AddVoucherCategory handles adding a category to a voucher
func (h *Voucher) AddVoucherCategory(w http.ResponseWriter, r *http.Request) {
	voucherID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	categoryID, err := uuid.Parse(chi.URLParam(r, "category_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid category ID"))
		return
	}

	if err := h.service.AddVoucherCategory(r.Context(), voucherID, categoryID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}

// RemoveVoucherCategory handles removing a category from a voucher
func (h *Voucher) RemoveVoucherCategory(w http.ResponseWriter, r *http.Request) {
	voucherID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid discount ID"))
		return
	}

	categoryID, err := uuid.Parse(chi.URLParam(r, "category_id"))
	if err != nil {
		commonjson.RespondBadRequestError(w, errors.New("invalid category ID"))
		return
	}

	if err := h.service.RemoveVoucherCategory(r.Context(), voucherID, categoryID); err != nil {
		commonjson.RespondInternalServerError(w, err)
		return
	}

	commonjson.RespondNoContent(w)
}
