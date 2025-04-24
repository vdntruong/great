package services

import (
	"context"
	"errors"

	"oop/internal/domain"
	"oop/internal/repositories"
)

// ProductService provides business logic for products
// Single Responsibility Principle: Only handles product-related business logic
type ProductService struct {
	productRepo repositories.ProductRepository
}

// NewProductService creates a new product service
// Dependency Inversion Principle: Depends on the repository interface, not implementation
func NewProductService(productRepo repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// GetProductByID gets a product by ID
func (s *ProductService) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	return s.productRepo.FindByID(ctx, id)
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) error {
	// Validate product before saving
	if err := product.Validate(); err != nil {
		return err
	}

	return s.productRepo.Save(ctx, product)
}

// UpdateProductInventory updates a product's inventory
func (s *ProductService) UpdateProductInventory(ctx context.Context, id string, change int) error {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Check if we have enough inventory
	newInventory := product.Inventory + change
	if newInventory < 0 {
		return errors.New("insufficient inventory")
	}

	product.Inventory = newInventory
	return s.productRepo.Update(ctx, product)
}
