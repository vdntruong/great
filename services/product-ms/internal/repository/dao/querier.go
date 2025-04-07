// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package dao

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddDiscountCategory(ctx context.Context, arg *AddDiscountCategoryParams) error
	AddDiscountProduct(ctx context.Context, arg *AddDiscountProductParams) error
	AddVoucherCategory(ctx context.Context, arg *AddVoucherCategoryParams) error
	AddVoucherProduct(ctx context.Context, arg *AddVoucherProductParams) error
	CountDiscounts(ctx context.Context, storeID uuid.UUID) (int64, error)
	CountStores(ctx context.Context) (int64, error)
	CreateDiscount(ctx context.Context, arg *CreateDiscountParams) (*Discount, error)
	CreateProduct(ctx context.Context, arg *CreateProductParams) (*Product, error)
	CreateProductImage(ctx context.Context, arg *CreateProductImageParams) (*ProductImage, error)
	CreateProductVariant(ctx context.Context, arg *CreateProductVariantParams) (*ProductVariant, error)
	CreateStore(ctx context.Context, arg *CreateStoreParams) (*Store, error)
	CreateStoreCategory(ctx context.Context, arg *CreateStoreCategoryParams) (*StoreCategory, error)
	CreateVoucher(ctx context.Context, arg *CreateVoucherParams) (*Voucher, error)
	DeleteDiscount(ctx context.Context, id uuid.UUID) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProductImage(ctx context.Context, id uuid.UUID) error
	DeleteProductVariant(ctx context.Context, id uuid.UUID) error
	DeleteStore(ctx context.Context, id uuid.UUID) error
	DeleteStoreCategory(ctx context.Context, id uuid.UUID) error
	DeleteVoucher(ctx context.Context, id uuid.UUID) error
	GetDiscountByCode(ctx context.Context, arg *GetDiscountByCodeParams) (*Discount, error)
	GetDiscountByID(ctx context.Context, id uuid.UUID) (*Discount, error)
	GetDiscountCategories(ctx context.Context, discountID uuid.UUID) ([]*StoreCategory, error)
	GetDiscountProducts(ctx context.Context, discountID uuid.UUID) ([]*Product, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*Product, error)
	GetProductBySlug(ctx context.Context, arg *GetProductBySlugParams) (*Product, error)
	GetProductImageByID(ctx context.Context, id uuid.UUID) (*ProductImage, error)
	GetProductImages(ctx context.Context, productID uuid.UUID) ([]*ProductImage, error)
	GetProductVariantByID(ctx context.Context, id uuid.UUID) (*ProductVariant, error)
	GetProductVariants(ctx context.Context, productID uuid.UUID) ([]*ProductVariant, error)
	GetStoreByID(ctx context.Context, id uuid.UUID) (*Store, error)
	GetStoreBySlug(ctx context.Context, slug string) (*Store, error)
	GetStoreCategories(ctx context.Context, storeID uuid.UUID) ([]*StoreCategory, error)
	GetStoreCategoryByID(ctx context.Context, id uuid.UUID) (*StoreCategory, error)
	GetStoreCategoryBySlug(ctx context.Context, arg *GetStoreCategoryBySlugParams) (*StoreCategory, error)
	GetVoucherByCode(ctx context.Context, arg *GetVoucherByCodeParams) (*Voucher, error)
	GetVoucherByID(ctx context.Context, id uuid.UUID) (*Voucher, error)
	IncrementVoucherUsage(ctx context.Context, id uuid.UUID) error
	ListDiscounts(ctx context.Context, arg *ListDiscountsParams) ([]*Discount, error)
	ListProducts(ctx context.Context, arg *ListProductsParams) ([]*Product, error)
	ListProductsByCategory(ctx context.Context, arg *ListProductsByCategoryParams) ([]*Product, error)
	ListStores(ctx context.Context, arg *ListStoresParams) ([]*Store, error)
	ListVouchers(ctx context.Context, arg *ListVouchersParams) ([]*Voucher, error)
	RemoveDiscountCategory(ctx context.Context, arg *RemoveDiscountCategoryParams) error
	RemoveDiscountProduct(ctx context.Context, arg *RemoveDiscountProductParams) error
	RemoveVoucherCategory(ctx context.Context, arg *RemoveVoucherCategoryParams) error
	RemoveVoucherProduct(ctx context.Context, arg *RemoveVoucherProductParams) error
	UpdateDiscount(ctx context.Context, arg *UpdateDiscountParams) (*Discount, error)
	UpdateProduct(ctx context.Context, arg *UpdateProductParams) (*Product, error)
	UpdateProductImage(ctx context.Context, arg *UpdateProductImageParams) (*ProductImage, error)
	UpdateProductVariant(ctx context.Context, arg *UpdateProductVariantParams) (*ProductVariant, error)
	UpdateStore(ctx context.Context, arg *UpdateStoreParams) (*Store, error)
	UpdateStoreCategory(ctx context.Context, arg *UpdateStoreCategoryParams) (*StoreCategory, error)
	UpdateVoucher(ctx context.Context, arg *UpdateVoucherParams) (*Voucher, error)
	UpdateVoucherStatus(ctx context.Context, arg *UpdateVoucherStatusParams) error
}

var _ Querier = (*Queries)(nil)
