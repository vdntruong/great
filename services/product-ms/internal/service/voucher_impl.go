package service

import (
	"context"

	"product-ms/internal/models"
	"product-ms/internal/repository/dao"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// VoucherServiceImpl implements VoucherService
type VoucherServiceImpl struct {
	dao       *dao.Queries
	validator validator.VoucherValidator
}

var _ VoucherService = (*VoucherServiceImpl)(nil)

// NewVoucherService creates a new VoucherService
func NewVoucherService(dao *dao.Queries) *VoucherServiceImpl {
	return &VoucherServiceImpl{
		dao:       dao,
		validator: validator.NewVoucherValidator(),
	}
}

// ToVoucherModel converts a DAO voucher to a model voucher
func ToVoucherModel(v *dao.Voucher) *models.Voucher {
	if v == nil {
		return nil
	}

	return &models.Voucher{
		ID:                v.ID,
		StoreID:           v.StoreID,
		Code:              v.Code,
		Type:              models.VoucherType(v.Type),
		Value:             stringToFloat64Ptr(v.Value),
		MinPurchaseAmount: stringToFloat64Ptr(v.MinPurchaseAmount),
		MaxDiscountAmount: stringToFloat64Ptr(v.MaxDiscountAmount),
		StartDate:         v.StartDate,
		EndDate:           nullTimeToTimePtr(v.EndDate),
		UsageLimit:        nullInt32ToInt32Ptr(v.UsageLimit),
		UsageCount:        nullInt32ToInt32(v.UsageCount),
		Status:            models.VoucherStatus(v.Status),
		CreatedAt:         v.CreatedAt.Time,
		UpdatedAt:         v.UpdatedAt.Time,
	}
}

// ToVoucherListModel converts a DAO voucher list to a model voucher list
func ToVoucherListModel(vouchers []*dao.Voucher, total int64) *models.VoucherList {
	items := make([]models.Voucher, len(vouchers))
	for i, v := range vouchers {
		items[i] = *ToVoucherModel(v)
	}

	return &models.VoucherList{
		Vouchers:   items,
		TotalCount: total,
		Page:       1, // TODO: Calculate page based on offset and limit
		Limit:      10, // TODO: Use actual limit
	}
}

// CreateVoucher creates a new voucher
func (s *VoucherServiceImpl) CreateVoucher(ctx context.Context, params models.CreateVoucherParams) (*models.Voucher, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	id := uuid.New()
	daoParams := dao.CreateVoucherParams{
		ID:                id,
		StoreID:           params.StoreID,
		Code:              params.Code,
		Type:              dao.VoucherType(params.Type),
		Value:             float64PtrToString(params.Value),
		MinPurchaseAmount: float64PtrToString(params.MinPurchaseAmount),
		MaxDiscountAmount: float64PtrToString(params.MaxDiscountAmount),
		StartDate:         params.StartDate,
		EndDate:           timePtrToNullTime(params.EndDate),
		UsageLimit:        int32PtrToNullInt32(params.UsageLimit),
		Status:            dao.VoucherStatus(params.Status),
	}

	voucher, err := s.dao.CreateVoucher(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToVoucherModel(voucher), nil
}

// GetVoucherByID gets a voucher by ID
func (s *VoucherServiceImpl) GetVoucherByID(ctx context.Context, id string) (*models.Voucher, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	voucher, err := s.dao.GetVoucherByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return ToVoucherModel(voucher), nil
}

// ListVouchers lists vouchers
func (s *VoucherServiceImpl) ListVouchers(ctx context.Context, params models.ListVouchersParams) (*models.VoucherList, error) {
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	daoParams := dao.ListVouchersParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}

	vouchers, err := s.dao.ListVouchers(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToVoucherListModel(vouchers, int64(len(vouchers))), nil
}

// UpdateVoucher updates a voucher
func (s *VoucherServiceImpl) UpdateVoucher(ctx context.Context, params models.UpdateVoucherParams) (*models.Voucher, error) {
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	daoParams := dao.UpdateVoucherParams{
		ID:         params.ID,
		Column2:    params.Code,
		Column3:    (*string)(params.Type),
		Column4:    float64PtrToString(params.Value),
		Column5:    float64PtrToString(params.MinPurchaseAmount),
		Column6:    float64PtrToString(params.MaxDiscountAmount),
		StartDate:  *params.StartDate,
		EndDate:    timePtrToNullTime(params.EndDate),
		UsageLimit: int32PtrToNullInt32(params.UsageLimit),
		Column10:   (*string)(params.Status),
	}

	voucher, err := s.dao.UpdateVoucher(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ToVoucherModel(voucher), nil
}

// DeleteVoucher deletes a voucher
func (s *VoucherServiceImpl) DeleteVoucher(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.dao.DeleteVoucher(ctx, uuid)
}

// UpdateVoucherStatus updates a voucher's status
func (s *VoucherServiceImpl) UpdateVoucherStatus(ctx context.Context, id uuid.UUID, status models.VoucherStatus) error {
	params := dao.UpdateVoucherStatusParams{
		ID:     id,
		Status: dao.VoucherStatus(status),
	}
	return s.dao.UpdateVoucherStatus(ctx, &params)
}

// AddVoucherProduct adds a product to a voucher
func (s *VoucherServiceImpl) AddVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao.AddVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.dao.AddVoucherProduct(ctx, &params)
}

// RemoveVoucherProduct removes a product from a voucher
func (s *VoucherServiceImpl) RemoveVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao.RemoveVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.dao.RemoveVoucherProduct(ctx, &params)
}

// AddVoucherCategory adds a category to a voucher
func (s *VoucherServiceImpl) AddVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao.AddVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.dao.AddVoucherCategory(ctx, &params)
}

// RemoveVoucherCategory removes a category from a voucher
func (s *VoucherServiceImpl) RemoveVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao.RemoveVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.dao.RemoveVoucherCategory(ctx, &params)
}
