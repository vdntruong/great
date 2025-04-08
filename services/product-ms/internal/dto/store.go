package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateStoreRequest represents the request body for creating a store
type CreateStoreRequest struct {
	Name         string                 `json:"name"`
	Slug         string                 `json:"slug"`
	Description  string                 `json:"description"`
	LogoURL      string                 `json:"logo_url"`
	CoverURL     string                 `json:"cover_url"`
	Status       string                 `json:"status"`
	IsVerified   bool                   `json:"is_verified"`
	OwnerID      uuid.UUID              `json:"owner_id"`
	ContactEmail string                 `json:"contact_email"`
	ContactPhone string                 `json:"contact_phone"`
	Address      string                 `json:"address"`
	Settings     map[string]interface{} `json:"settings"`
}

// UpdateStoreRequest represents the request body for updating a store
type UpdateStoreRequest struct {
	Name         *string                 `json:"name,omitempty"`
	Slug         *string                 `json:"slug,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	LogoURL      *string                 `json:"logo_url,omitempty"`
	CoverURL     *string                 `json:"cover_url,omitempty"`
	Status       *string                 `json:"status,omitempty"`
	IsVerified   *bool                   `json:"is_verified,omitempty"`
	OwnerID      *uuid.UUID              `json:"owner_id,omitempty"`
	ContactEmail *string                 `json:"contact_email,omitempty"`
	ContactPhone *string                 `json:"contact_phone,omitempty"`
	Address      *string                 `json:"address,omitempty"`
	Settings     *map[string]interface{} `json:"settings,omitempty"`
}

// ListStoresRequest represents the query parameters for listing stores
type ListStoresRequest struct {
	Page      int    `form:"page" validate:"min=1"`
	Limit     int    `form:"limit" validate:"min=1,max=100"`
	SortBy    string `form:"sort_by" validate:"oneof=name slug status created_at updated_at"`
	SortOrder string `form:"sort_order" validate:"oneof=asc desc"`
	Status    string `form:"status" validate:"omitempty,oneof=active inactive"`
}

// StoreResponse represents the response for a store
type StoreResponse struct {
	ID           uuid.UUID              `json:"id"`
	Name         string                 `json:"name"`
	Slug         string                 `json:"slug"`
	Description  string                 `json:"description"`
	LogoURL      string                 `json:"logo_url"`
	CoverURL     string                 `json:"cover_url"`
	Status       string                 `json:"status"`
	IsVerified   bool                   `json:"is_verified"`
	OwnerID      uuid.UUID              `json:"owner_id"`
	ContactEmail string                 `json:"contact_email"`
	ContactPhone string                 `json:"contact_phone"`
	Address      string                 `json:"address"`
	Settings     map[string]interface{} `json:"settings"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

// StoreListResponse represents the response for listing stores
type StoreListResponse struct {
	Stores     []StoreResponse `json:"stores"`
	TotalCount int64           `json:"total_count"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
}
