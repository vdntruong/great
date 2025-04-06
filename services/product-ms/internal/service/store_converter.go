package service

import (
	"database/sql"
	"encoding/json"
	"product-ms/internal/models"
	"product-ms/internal/repository/dao"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

// ConvertStoreToModel converts a DAO store to a model store
func ConvertStoreToModel(store *dao.Store) *models.Store {
	if store == nil {
		return nil
	}

	var settings map[string]interface{}
	if store.Settings.Valid {
		_ = json.Unmarshal(store.Settings.RawMessage, &settings)
	}

	return &models.Store{
		ID:           store.ID,
		Name:         store.Name,
		Slug:         store.Slug,
		Description:  store.Description.String,
		LogoURL:      store.LogoUrl.String,
		CoverURL:     store.CoverUrl.String,
		Status:       string(store.Status),
		IsVerified:   store.IsVerified,
		OwnerID:      store.OwnerID,
		ContactEmail: store.ContactEmail.String,
		ContactPhone: store.ContactPhone.String,
		Address:      store.Address.String,
		Settings:     settings,
		CreatedAt:    store.CreatedAt.Time,
		UpdatedAt:    store.UpdatedAt.Time,
	}
}

// ConvertStoreListToModel converts a DAO store list to a model store list
func ConvertStoreListToModel(stores []*dao.Store, totalCount int64, page, limit int) *models.StoreList {
	modelStores := make([]models.Store, len(stores))
	for i, store := range stores {
		modelStores[i] = *ConvertStoreToModel(store)
	}

	return &models.StoreList{
		Stores:     modelStores,
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
	}
}

// ConvertCreateStoreParamsToDAO converts a model CreateStoreParams to DAO CreateStoreParams
func ConvertCreateStoreParamsToDAO(params models.CreateStoreParams) *dao.CreateStoreParams {
	settings := pqtype.NullRawMessage{}
	if params.Settings != nil {
		raw, _ := json.Marshal(params.Settings)
		settings.RawMessage = raw
		settings.Valid = true
	}

	return &dao.CreateStoreParams{
		Name:         params.Name,
		Slug:         params.Slug,
		Description:  sql.NullString{String: params.Description, Valid: params.Description != ""},
		LogoUrl:      sql.NullString{String: params.LogoURL, Valid: params.LogoURL != ""},
		CoverUrl:     sql.NullString{String: params.CoverURL, Valid: params.CoverURL != ""},
		Status:       dao.StoreStatus(params.Status),
		IsVerified:   params.IsVerified,
		OwnerID:      params.OwnerID,
		ContactEmail: sql.NullString{String: params.ContactEmail, Valid: params.ContactEmail != ""},
		ContactPhone: sql.NullString{String: params.ContactPhone, Valid: params.ContactPhone != ""},
		Address:      sql.NullString{String: params.Address, Valid: params.Address != ""},
		Settings:     settings,
	}
}

// ConvertUpdateStoreParamsToDAO converts a model UpdateStoreParams to DAO UpdateStoreParams
func ConvertUpdateStoreParamsToDAO(id uuid.UUID, params models.UpdateStoreParams) *dao.UpdateStoreParams {
	daoParams := &dao.UpdateStoreParams{
		ID: id,
	}

	if params.Name != "" {
		daoParams.Name = params.Name
	}
	if params.Slug != "" {
		daoParams.Slug = params.Slug
	}
	if params.Description != "" {
		daoParams.Description = sql.NullString{String: params.Description, Valid: true}
	}
	if params.LogoURL != "" {
		daoParams.LogoUrl = sql.NullString{String: params.LogoURL, Valid: true}
	}
	if params.CoverURL != "" {
		daoParams.CoverUrl = sql.NullString{String: params.CoverURL, Valid: true}
	}
	if params.Status != "" {
		daoParams.Status = dao.StoreStatus(params.Status)
	}
	daoParams.IsVerified = params.IsVerified
	if params.ContactEmail != "" {
		daoParams.ContactEmail = sql.NullString{String: params.ContactEmail, Valid: true}
	}
	if params.ContactPhone != "" {
		daoParams.ContactPhone = sql.NullString{String: params.ContactPhone, Valid: true}
	}
	if params.Address != "" {
		daoParams.Address = sql.NullString{String: params.Address, Valid: true}
	}
	if params.Settings != nil {
		raw, _ := json.Marshal(params.Settings)
		daoParams.Settings = pqtype.NullRawMessage{RawMessage: raw, Valid: true}
	}

	return daoParams
}

// ConvertListStoresParamsToDAO converts a model ListStoresParams to DAO ListStoresParams
func ConvertListStoresParamsToDAO(params models.ListStoresParams) *dao.ListStoresParams {
	return &dao.ListStoresParams{
		Limit:  int32(params.Limit),
		Offset: int32((params.Page - 1) * params.Limit),
	}
}
