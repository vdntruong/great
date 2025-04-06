package service

import (
	"fmt"
	"product-ms/internal/models"
	"product-ms/internal/repository/dao"
	"regexp"
)

// ValidateCreateStoreParams validates the parameters for creating a store
func ValidateCreateStoreParams(params models.CreateStoreParams) error {
	if params.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(params.Name) > 100 {
		return fmt.Errorf("name must be less than 100 characters")
	}

	if params.Slug == "" {
		return fmt.Errorf("slug is required")
	}
	if len(params.Slug) > 100 {
		return fmt.Errorf("slug must be less than 100 characters")
	}
	if !isValidSlug(params.Slug) {
		return fmt.Errorf("slug must contain only lowercase letters, numbers, and hyphens")
	}

	if len(params.Description) > 500 {
		return fmt.Errorf("description must be less than 500 characters")
	}

	if !isValidStoreStatus(params.Status) {
		return fmt.Errorf("invalid status: %s. Valid statuses are: pending, active, suspended, closed", params.Status)
	}

	return nil
}

// ValidateUpdateStoreParams validates the parameters for updating a store
func ValidateUpdateStoreParams(params models.UpdateStoreParams) error {
	if params.Name != "" && len(params.Name) > 100 {
		return fmt.Errorf("name must be less than 100 characters")
	}

	if params.Slug != "" {
		if len(params.Slug) > 100 {
			return fmt.Errorf("slug must be less than 100 characters")
		}
		if !isValidSlug(params.Slug) {
			return fmt.Errorf("slug must contain only lowercase letters, numbers, and hyphens")
		}
	}

	if params.Description != "" && len(params.Description) > 500 {
		return fmt.Errorf("description must be less than 500 characters")
	}

	if params.Status != "" && !isValidStoreStatus(params.Status) {
		return fmt.Errorf("invalid status: %s. Valid statuses are: pending, active, suspended, closed", params.Status)
	}

	return nil
}

// isValidSlug checks if a string is a valid slug
func isValidSlug(slug string) bool {
	// Slug should only contain lowercase letters, numbers, and hyphens
	// Should not start or end with a hyphen
	// Should not contain consecutive hyphens
	slugRegex := regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	return slugRegex.MatchString(slug)
}

// isValidStoreStatus checks if a status is valid
func isValidStoreStatus(status string) bool {
	switch dao.StoreStatus(status) {
	case dao.StoreStatusPending, dao.StoreStatusActive, dao.StoreStatusSuspended, dao.StoreStatusClosed:
		return true
	default:
		return false
	}
}
