package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"product-ms/internal/models"
	"product-ms/internal/repository/dao"

	"github.com/google/uuid"
)

// storeService implements the StoreService interface
type storeService struct {
	storeDAO *dao.Queries
}

var _ StoreService = (*storeService)(nil)

// NewStoreService creates a new instance of StoreService
func NewStoreService(storeDAO *dao.Queries) StoreService {
	return &storeService{
		storeDAO: storeDAO,
	}
}

func (s *storeService) CreateStore(ctx context.Context, params models.CreateStoreParams) (*models.Store, error) {
	// Validate input parameters
	if err := ValidateCreateStoreParams(params); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}

	// Convert and create store in database
	daoParams := ConvertCreateStoreParamsToDAO(params)
	store, err := s.storeDAO.CreateStore(ctx, daoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

func (s *storeService) GetStoreByID(ctx context.Context, id string) (*models.Store, error) {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid store ID: %w", err)
	}

	store, err := s.storeDAO.GetStoreByID(ctx, storeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("store not found")
		}
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

func (s *storeService) ListStores(ctx context.Context, params models.ListStoresParams) (*models.StoreList, error) {
	// Convert and get stores
	daoParams := ConvertListStoresParamsToDAO(params)
	stores, err := s.storeDAO.ListStores(ctx, daoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to list stores: %w", err)
	}

	// For now, use the length of stores as total count
	// TODO: Implement proper counting with a separate query
	totalCount := int64(len(stores))

	return ConvertStoreListToModel(stores, totalCount, params.Page, params.Limit), nil
}

func (s *storeService) UpdateStore(ctx context.Context, id string, params models.UpdateStoreParams) (*models.Store, error) {
	// Validate input parameters
	if err := ValidateUpdateStoreParams(params); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}

	storeID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid store ID: %w", err)
	}

	// Convert and update store in database
	daoParams := ConvertUpdateStoreParamsToDAO(storeID, params)
	store, err := s.storeDAO.UpdateStore(ctx, daoParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("store not found")
		}
		return nil, fmt.Errorf("failed to update store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

func (s *storeService) DeleteStore(ctx context.Context, id string) error {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid store ID: %w", err)
	}

	err = s.storeDAO.DeleteStore(ctx, storeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("store not found")
		}
		return fmt.Errorf("failed to delete store: %w", err)
	}

	return nil
}
