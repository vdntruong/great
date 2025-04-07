package service

import (
	"context"

	"product-ms/db/dao"
	"product-ms/internal/models"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// VoucherServiceImpl implements VoucherService
type VoucherServiceImpl struct {
	queries   *dao.Queries
	validator validator.VoucherValidator
}

var _ VoucherService = (*VoucherServiceImpl)(nil)

// NewVoucherService creates a new VoucherService
func NewVoucherService(queries *dao.Queries) *VoucherServiceImpl {
	return &VoucherServiceImpl{
		queries:   queries,
		validator: validator.NewVoucherValidator(),
	}
}

// CreateVoucher creates a new voucher
func (s *VoucherServiceImpl) CreateVoucher(ctx context.Context, params models.CreateVoucherParams) (*models.Voucher, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertCreateVoucherParamsToDAO(params)
	voucher, err := s.queries.CreateVoucher(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertVoucherToModel(voucher), nil
}

// GetVoucherByID gets a voucher by ID
func (s *VoucherServiceImpl) GetVoucherByID(ctx context.Context, id string) (*models.Voucher, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	voucher, err := s.queries.GetVoucher(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return ConvertVoucherToModel(voucher), nil
}

// ListVouchers lists vouchers
func (s *VoucherServiceImpl) ListVouchers(ctx context.Context, params models.ListVouchersParams) (*models.VoucherList, error) {
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	daoParams := ConvertListVouchersParamsToDAO(params)
	vouchers, err := s.queries.ListVouchers(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	total, err := s.queries.CountVouchers(ctx, params.StoreID)
	if err != nil {
		return nil, err
	}

	page := int(params.Offset/params.Limit) + 1
	return ConvertVoucherListToModel(vouchers, total, page, int(params.Limit)), nil
}

// UpdateVoucher updates a voucher
func (s *VoucherServiceImpl) UpdateVoucher(ctx context.Context, params models.UpdateVoucherParams) (*models.Voucher, error) {
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertUpdateVoucherParamsToDAO(params)
	voucher, err := s.queries.UpdateVoucher(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertVoucherToModel(voucher), nil
}

// DeleteVoucher deletes a voucher
func (s *VoucherServiceImpl) DeleteVoucher(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.queries.DeleteVoucher(ctx, uuid)
}

// AddVoucherProduct adds a product to a voucher
func (s *VoucherServiceImpl) AddVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao.AddVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.queries.AddVoucherProduct(ctx, &params)
}

// RemoveVoucherProduct removes a product from a voucher
func (s *VoucherServiceImpl) RemoveVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao.RemoveVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.queries.RemoveVoucherProduct(ctx, &params)
}

// AddVoucherCategory adds a category to a voucher
func (s *VoucherServiceImpl) AddVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao.AddVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.queries.AddVoucherCategory(ctx, &params)
}

// RemoveVoucherCategory removes a category from a voucher
func (s *VoucherServiceImpl) RemoveVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao.RemoveVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.queries.RemoveVoucherCategory(ctx, &params)
}
