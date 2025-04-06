package validator

import (
	"errors"
	"fmt"
	"strings"

	"product-ms/internal/models"

	"github.com/go-playground/validator/v10"
)

// ProductValidator defines the interface for product validation
type ProductValidator interface {
	ValidateCreate(params models.CreateProductParams) error
	ValidateUpdate(params models.UpdateProductParams) error
	ValidateList(params models.ListProductsParams) error
}

// productValidator implements ProductValidator
type productValidator struct {
	validate *validator.Validate
}

// NewProductValidator creates a new ProductValidator
func NewProductValidator() ProductValidator {
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
	return &productValidator{validate: v}
}

// ValidateCreate validates product creation parameters
func (v *productValidator) ValidateCreate(params models.CreateProductParams) error {
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
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateUpdate validates product update parameters
func (v *productValidator) ValidateUpdate(params models.UpdateProductParams) error {
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
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateList validates product listing parameters
func (v *productValidator) ValidateList(params models.ListProductsParams) error {
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
