package service

import (
	"context"
	dao2 "product-ms/db/dao"

	"product-ms/internal/models"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// VoucherServiceImpl implements VoucherService
type VoucherServiceImpl struct {
	dao       *dao2.Queries
	validator validator.VoucherValidator
}

var _ VoucherService = (*VoucherServiceImpl)(nil)

// NewVoucherService creates a new VoucherService
func NewVoucherService(dao *dao2.Queries) *VoucherServiceImpl {
	return &VoucherServiceImpl{
		dao:       dao,
		validator: validator.NewVoucherValidator(),
	}
}

// CreateVoucher creates a new voucher
func (s *VoucherServiceImpl) CreateVoucher(ctx context.Context, params models.CreateVoucherParams) (*models.Voucher, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertCreateVoucherParamsToDAO(params)
	voucher, err := s.dao.CreateVoucher(ctx, &daoParams)
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

	voucher, err := s.dao.GetVoucher(ctx, uuid)
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
	vouchers, err := s.dao.ListVouchers(ctx, &daoParams)
	if err != nil {
		return nil, err
	}

	total, err := s.dao.CountVouchers(ctx, params.StoreID)
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
	voucher, err := s.dao.UpdateVoucher(ctx, &daoParams)
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
	return s.dao.DeleteVoucher(ctx, uuid)
}

// AddVoucherProduct adds a product to a voucher
func (s *VoucherServiceImpl) AddVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao2.AddVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.dao.AddVoucherProduct(ctx, &params)
}

// RemoveVoucherProduct removes a product from a voucher
func (s *VoucherServiceImpl) RemoveVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error {
	params := dao2.RemoveVoucherProductParams{
		VoucherID: voucherID,
		ProductID: productID,
	}
	return s.dao.RemoveVoucherProduct(ctx, &params)
}

// AddVoucherCategory adds a category to a voucher
func (s *VoucherServiceImpl) AddVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao2.AddVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.dao.AddVoucherCategory(ctx, &params)
}

// RemoveVoucherCategory removes a category from a voucher
func (s *VoucherServiceImpl) RemoveVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error {
	params := dao2.RemoveVoucherCategoryParams{
		VoucherID:  voucherID,
		CategoryID: categoryID,
	}
	return s.dao.RemoveVoucherCategory(ctx, &params)
}
