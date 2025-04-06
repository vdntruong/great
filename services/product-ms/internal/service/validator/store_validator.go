package validator

import (
	"errors"
	"fmt"
	"strings"

	"product-ms/internal/models"

	"github.com/go-playground/validator/v10"
)

// StoreValidator defines the interface for store validation
type StoreValidator interface {
	ValidateCreate(params models.CreateStoreParams) error
	ValidateUpdate(params models.UpdateStoreParams) error
	ValidateList(params models.ListStoresParams) error
}

// storeValidator implements StoreValidator
type storeValidator struct {
	validate *validator.Validate
}

// NewStoreValidator creates a new StoreValidator
func NewStoreValidator() StoreValidator {
	v := validator.New()
	// Register custom validation for slug format
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		slug := fl.Field().String()
		if slug == "" {
			return false
		}
		// Slug should only contain lowercase letters, numbers, and hyphens
		for _, c := range slug {
			if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-') {
				return false
			}
		}
		return true
	})
	return &storeValidator{validate: v}
}

// ValidateCreate validates store creation parameters
func (v *storeValidator) ValidateCreate(params models.CreateStoreParams) error {
	if err := v.validate.Struct(params); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				switch e.Tag() {
				case "required":
					return fmt.Errorf("%s is required", strings.ToLower(e.Field()))
				case "min":
					return fmt.Errorf("%s must be at least %s", strings.ToLower(e.Field()), e.Param())
				case "max":
					return fmt.Errorf("%s must be at most %s characters", strings.ToLower(e.Field()), e.Param())
				case "oneof":
					return fmt.Errorf("%s must be one of: %s", strings.ToLower(e.Field()), e.Param())
				case "slug":
					return fmt.Errorf("%s must be a valid slug (lowercase letters, numbers, and hyphens only)", strings.ToLower(e.Field()))
				case "email":
					return fmt.Errorf("%s must be a valid email address", strings.ToLower(e.Field()))
				case "url":
					return fmt.Errorf("%s must be a valid URL", strings.ToLower(e.Field()))
				case "e164":
					return fmt.Errorf("%s must be a valid phone number in E.164 format", strings.ToLower(e.Field()))
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateUpdate validates store update parameters
func (v *storeValidator) ValidateUpdate(params models.UpdateStoreParams) error {
	if err := v.validate.Struct(params); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				switch e.Tag() {
				case "required":
					return fmt.Errorf("%s is required", strings.ToLower(e.Field()))
				case "min":
					return fmt.Errorf("%s must be at least %s", strings.ToLower(e.Field()), e.Param())
				case "max":
					return fmt.Errorf("%s must be at most %s characters", strings.ToLower(e.Field()), e.Param())
				case "oneof":
					return fmt.Errorf("%s must be one of: %s", strings.ToLower(e.Field()), e.Param())
				case "slug":
					return fmt.Errorf("%s must be a valid slug (lowercase letters, numbers, and hyphens only)", strings.ToLower(e.Field()))
				case "email":
					return fmt.Errorf("%s must be a valid email address", strings.ToLower(e.Field()))
				case "url":
					return fmt.Errorf("%s must be a valid URL", strings.ToLower(e.Field()))
				case "e164":
					return fmt.Errorf("%s must be a valid phone number in E.164 format", strings.ToLower(e.Field()))
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateList validates store listing parameters
func (v *storeValidator) ValidateList(params models.ListStoresParams) error {
	if err := v.validate.Struct(params); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				switch e.Tag() {
				case "min":
					return fmt.Errorf("%s must be at least %s", strings.ToLower(e.Field()), e.Param())
				case "max":
					return fmt.Errorf("%s must be at most %s", strings.ToLower(e.Field()), e.Param())
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}
