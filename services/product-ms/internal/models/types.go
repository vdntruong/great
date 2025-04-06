package models

import "github.com/google/uuid"

// ProductID represents a product's unique identifier
type ProductID string

// NewProductID creates a new ProductID from a string
func NewProductID(id string) (ProductID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", err
	}
	return ProductID(id), nil
}

// StoreID represents a store's unique identifier
type StoreID string

// NewStoreID creates a new StoreID from a string
func NewStoreID(id string) (StoreID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", err
	}
	return StoreID(id), nil
}
