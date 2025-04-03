package repository

import (
	"database/sql"

	"product-ms/internal/dao"
)

type ProductRepo struct {
	*sql.DB
	DAO *dao.Queries
}

func NewProductRepository(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		DB:  db,
		DAO: dao.New(db),
	}
}
