package service

import (
	"context"
	"fmt"
	"product-ms/db/dao"

	"product-ms/internal/models"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	queries   *dao.Queries
	validator validator.ProductValidator
}

var _ ProductService = (*ProductServiceImpl)(nil)

func NewProductService(queries *dao.Queries) ProductService {
	return &ProductServiceImpl{
		queries:   queries,
		validator: validator.NewProductValidator(),
	}
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, params models.CreateProductParams) (*models.Product, error) {
	// Validate input parameters
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// Convert params to DAO params
	daoParams, err := ConvertCreateProductParamsToDAO(params)
	if err != nil {
		return nil, err
	}

	// Create product in database
	product, err := s.queries.CreateProduct(ctx, daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertProductToModel(product), nil
}

func (s *ProductServiceImpl) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	productID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID: %w", err)
	}

	product, err := s.queries.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	return ConvertProductToModel(product), nil
}

func (s *ProductServiceImpl) ListProducts(ctx context.Context, params models.ListProductsParams) ([]*models.Product, error) {
	// Validate input parameters
	if err := s.validator.ValidateList(params); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	products, err := s.queries.ListProducts(ctx, ConvertListProductsParamsToDAO(params))
	if err != nil {
		return nil, err
	}

	result := make([]*models.Product, len(products))
	for i, product := range products {
		result[i] = ConvertProductToModel(product)
	}

	return result, nil
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, params models.UpdateProductParams) (*models.Product, error) {
	// Validate input parameters
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// Convert params to DAO params
	daoParams, err := ConvertUpdateProductParamsToDAO(params)
	if err != nil {
		return nil, err
	}

	// Update product in database
	product, err := s.queries.UpdateProduct(ctx, daoParams)
	if err != nil {
		return nil, err
	}

	return ConvertProductToModel(product), nil
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	productID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid product ID: %w", err)
	}

	return s.queries.DeleteProduct(ctx, productID)
}
