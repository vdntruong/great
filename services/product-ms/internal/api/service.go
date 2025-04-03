package api

import "product-ms/internal/repository"

type ProductSvc struct {
	ProductRepo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductSvc {
	return &ProductSvc{
		ProductRepo: productRepo,
	}
}
