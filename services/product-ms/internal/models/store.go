package models

import (
	"time"

	"github.com/google/uuid"
)

// Store represents a store in the system
type Store struct {
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

// CreateStoreParams represents the parameters for creating a new store
type CreateStoreParams struct {
	Name         string                 `json:"name" validate:"required,min=3,max=100"`
	Slug         string                 `json:"slug" validate:"required,min=3,max=100,slug"`
	Description  string                 `json:"description" validate:"max=500"`
	LogoURL      string                 `json:"logo_url"`
	CoverURL     string                 `json:"cover_url"`
	Status       string                 `json:"status" validate:"required,oneof=active inactive"`
	IsVerified   bool                   `json:"is_verified"`
	OwnerID      uuid.UUID              `json:"owner_id"`
	ContactEmail string                 `json:"contact_email"`
	ContactPhone string                 `json:"contact_phone"`
	Address      string                 `json:"address"`
	Settings     map[string]interface{} `json:"settings"`
}

// UpdateStoreParams represents the parameters for updating a store
type UpdateStoreParams struct {
	Name         string                 `json:"name" validate:"omitempty,min=3,max=100"`
	Slug         string                 `json:"slug" validate:"omitempty,min=3,max=100,slug"`
	Description  string                 `json:"description" validate:"omitempty,max=500"`
	LogoURL      string                 `json:"logo_url"`
	CoverURL     string                 `json:"cover_url"`
	Status       string                 `json:"status" validate:"omitempty,oneof=active inactive"`
	IsVerified   bool                   `json:"is_verified"`
	OwnerID      uuid.UUID              `json:"owner_id"`
	ContactEmail string                 `json:"contact_email"`
	ContactPhone string                 `json:"contact_phone"`
	Address      string                 `json:"address"`
	Settings     map[string]interface{} `json:"settings"`
}

// ListStoresParams represents the parameters for listing stores
type ListStoresParams struct {
	Page       int    `query:"page" validate:"min=1"`
	Limit      int    `query:"limit" validate:"min=1,max=100"`
	SortBy     string `query:"sort_by" validate:"oneof=name slug status created_at updated_at"`
	SortOrder  string `query:"sort_order" validate:"oneof=asc desc"`
	Status     string `query:"status" validate:"omitempty,oneof=active inactive"`
	IsVerified *bool  `query:"is_verified"`
}

// StoreList represents a list of stores with pagination info
type StoreList struct {
	Stores     []Store `json:"stores"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
}
