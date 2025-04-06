package service

import (
	"context"

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
