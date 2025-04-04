package service

import "product-ms/internal/repository/dao"

type ProductService struct {
	DAO dao.Querier
}

func NewProductService(productDAO dao.Querier) *ProductService {
	return &ProductService{
		DAO: productDAO,
	}
}
