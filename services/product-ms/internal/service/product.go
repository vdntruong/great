package service

import (
	"context"
	"github.com/google/uuid"
	"product-ms/internal/models"
)

// ProductService defines the interface for product operations
type ProductService interface {
	CreateProduct(ctx context.Context, params models.CreateProductParams) (*models.Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error)
	ListProducts(ctx context.Context, params models.ListProductsParams) ([]*models.Product, error)
	UpdateProduct(ctx context.Context, params models.UpdateProductParams) (*models.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
