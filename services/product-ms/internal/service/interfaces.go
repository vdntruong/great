package service

import (
	"context"
	"github.com/google/uuid"

	"product-ms/internal/models"
)

// StoreService defines the interface for store operations
type StoreService interface {
	CreateStore(ctx context.Context, params models.CreateStoreParams) (*models.Store, error)
	GetStoreByID(ctx context.Context, id string) (*models.Store, error)
	ListStores(ctx context.Context, params models.ListStoresParams) (*models.StoreList, error)
	UpdateStore(ctx context.Context, id string, params models.UpdateStoreParams) (*models.Store, error)
	DeleteStore(ctx context.Context, id string) error
}

// ProductService defines the interface for product operations
type ProductService interface {
	CreateProduct(ctx context.Context, params models.CreateProductParams) (*models.Product, error)
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	ListProducts(ctx context.Context, params models.ListProductsParams) ([]*models.Product, error)
	UpdateProduct(ctx context.Context, params models.UpdateProductParams) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

// DiscountService defines the interface for discount operations
type DiscountService interface {
	CreateDiscount(ctx context.Context, params models.CreateDiscountParams) (*models.Discount, error)
	GetDiscountByID(ctx context.Context, id string) (*models.Discount, error)
	ListDiscounts(ctx context.Context, params models.ListDiscountsParams) (*models.DiscountList, error)
	UpdateDiscount(ctx context.Context, params models.UpdateDiscountParams) (*models.Discount, error)
	DeleteDiscount(ctx context.Context, id string) error
	AddDiscountProduct(ctx context.Context, discountID, productID uuid.UUID) error
	RemoveDiscountProduct(ctx context.Context, discountID, productID uuid.UUID) error
	AddDiscountCategory(ctx context.Context, discountID, categoryID uuid.UUID) error
	RemoveDiscountCategory(ctx context.Context, discountID, categoryID uuid.UUID) error
}

// VoucherService defines the interface for voucher operations
type VoucherService interface {
	CreateVoucher(ctx context.Context, params models.CreateVoucherParams) (*models.Voucher, error)
	GetVoucherByID(ctx context.Context, id string) (*models.Voucher, error)
	ListVouchers(ctx context.Context, params models.ListVouchersParams) (*models.VoucherList, error)
	UpdateVoucher(ctx context.Context, params models.UpdateVoucherParams) (*models.Voucher, error)
	DeleteVoucher(ctx context.Context, id string) error
	AddVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error
	RemoveVoucherProduct(ctx context.Context, voucherID, productID uuid.UUID) error
	AddVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error
	RemoveVoucherCategory(ctx context.Context, voucherID, categoryID uuid.UUID) error
}
