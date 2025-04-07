package service

import (
	"context"
	"fmt"
	"product-ms/db/dao"

	"product-ms/internal/models"
	"product-ms/internal/service/validator"

	"github.com/google/uuid"
)

// StoreServiceImpl implements the StoreService interface
type StoreServiceImpl struct {
	queries   *dao.Queries
	validator validator.StoreValidator
}

var _ StoreService = (*StoreServiceImpl)(nil)

// NewStoreService creates a new store service
func NewStoreService(queries *dao.Queries) *StoreServiceImpl {
	return &StoreServiceImpl{
		queries:   queries,
		validator: validator.NewStoreValidator(),
	}
}

// CreateStore creates a new store
func (s *StoreServiceImpl) CreateStore(ctx context.Context, params models.CreateStoreParams) (*models.Store, error) {
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertCreateStoreParamsToDAO(params)
	store, err := s.queries.CreateStore(ctx, &daoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

// GetStoreByID retrieves a store by ID
func (s *StoreServiceImpl) GetStoreByID(ctx context.Context, id string) (*models.Store, error) {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid store ID: %w", err)
	}

	store, err := s.queries.GetStore(ctx, storeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

// GetStoreBySlug retrieves a store by slug
func (s *StoreServiceImpl) GetStoreBySlug(ctx context.Context, slug string) (*models.Store, error) {
	store, err := s.queries.GetStoreBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to get store by slug: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

// ListStores retrieves a list of stores
func (s *StoreServiceImpl) ListStores(ctx context.Context, params models.ListStoresParams) (*models.StoreList, error) {
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	daoParams := ConvertListStoresParamsToDAO(params)
	stores, err := s.queries.ListStores(ctx, &daoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to list stores: %w", err)
	}

	totalCount, err := s.queries.CountStores(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count stores: %w", err)
	}

	return ConvertStoreListToModel(stores, totalCount, int(params.Offset/params.Limit+1), int(params.Limit)), nil
}

// UpdateStore updates a store
func (s *StoreServiceImpl) UpdateStore(ctx context.Context, id string, params models.UpdateStoreParams) (*models.Store, error) {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid store ID: %w", err)
	}

	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	daoParams := ConvertUpdateStoreParamsToDAO(storeID, params)
	store, err := s.queries.UpdateStore(ctx, &daoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to update store: %w", err)
	}

	return ConvertStoreToModel(store), nil
}

// DeleteStore deletes a store
func (s *StoreServiceImpl) DeleteStore(ctx context.Context, id string) error {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid store ID: %w", err)
	}

	if err := s.queries.DeleteStore(ctx, storeID); err != nil {
		return fmt.Errorf("failed to delete store: %w", err)
	}

	return nil
}
