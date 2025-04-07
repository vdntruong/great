package validator

import (
	"errors"
	"fmt"
	"strings"

	"product-ms/internal/models"

	"github.com/go-playground/validator/v10"
)

// DiscountValidator defines the interface for discount validation
type DiscountValidator interface {
	ValidateCreate(params models.CreateDiscountParams) error
	ValidateUpdate(params models.UpdateDiscountParams) error
	ValidateList(params models.ListDiscountsParams) error
}

// discountValidator implements DiscountValidator
type discountValidator struct {
	validate *validator.Validate
}

// NewDiscountValidator creates a new DiscountValidator
func NewDiscountValidator() DiscountValidator {
	v := validator.New()
	return &discountValidator{validate: v}
}

// ValidateCreate validates discount creation parameters
func (v *discountValidator) ValidateCreate(params models.CreateDiscountParams) error {
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
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateUpdate validates discount update parameters
func (v *discountValidator) ValidateUpdate(params models.UpdateDiscountParams) error {
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
				default:
					return fmt.Errorf("invalid %s: %s", strings.ToLower(e.Field()), e.Tag())
				}
			}
		}
		return err
	}
	return nil
}

// ValidateList validates discount listing parameters
func (v *discountValidator) ValidateList(params models.ListDiscountsParams) error {
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
