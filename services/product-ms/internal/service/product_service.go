package service

import (
	"context"
	"product-ms/internal/models"
	"product-ms/internal/repository/dao"

	"github.com/google/uuid"
)

type productService struct {
	queries *dao.Queries
}

func NewProductService(queries *dao.Queries) ProductService {
	return &productService{
		queries: queries,
	}
}

func (s *productService) CreateProduct(ctx context.Context, params models.CreateProductParams) (*models.Product, error) {
	// Validate input parameters
	if err := ValidateCreateProductParams(params); err != nil {
		return nil, err
	}

	// Convert params to DAO params
	daoParams, err := convertCreateParams(params)
	if err != nil {
		return nil, err
	}

	// Generate new UUID for product ID
	daoParams.ID = uuid.New()

	// Create product in database
	product, err := s.queries.CreateProduct(ctx, daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertDAOProductToModel(product), nil
}

func (s *productService) GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := s.queries.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ConvertDAOProductToModel(product), nil
}

func (s *productService) ListProducts(ctx context.Context, params models.ListProductsParams) ([]*models.Product, error) {
	products, err := s.queries.ListProducts(ctx, &dao.ListProductsParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Product, len(products))
	for i, product := range products {
		result[i] = ConvertDAOProductToModel(product)
	}

	return result, nil
}

func (s *productService) UpdateProduct(ctx context.Context, params models.UpdateProductParams) (*models.Product, error) {
	// Validate input parameters
	if err := ValidateUpdateProductParams(params); err != nil {
		return nil, err
	}

	// Convert params to DAO params
	daoParams, err := convertUpdateParams(params)
	if err != nil {
		return nil, err
	}

	// Update product in database
	product, err := s.queries.UpdateProduct(ctx, daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertDAOProductToModel(product), nil
}

func (s *productService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteProduct(ctx, id)
}
