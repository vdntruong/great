package validator

import (
	"errors"
	"fmt"
	"strings"

	"product-ms/internal/models"

	"github.com/go-playground/validator/v10"
)

// VoucherValidator defines the interface for voucher validation
type VoucherValidator interface {
	ValidateCreate(params models.CreateVoucherParams) error
	ValidateUpdate(params models.UpdateVoucherParams) error
	ValidateList(params models.ListVouchersParams) error
}

// voucherValidator implements VoucherValidator
type voucherValidator struct {
	validate *validator.Validate
}

// NewVoucherValidator creates a new VoucherValidator
func NewVoucherValidator() VoucherValidator {
	v := validator.New()
	return &voucherValidator{validate: v}
}

// ValidateCreate validates voucher creation parameters
func (v *voucherValidator) ValidateCreate(params models.CreateVoucherParams) error {
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

// ValidateUpdate validates voucher update parameters
func (v *voucherValidator) ValidateUpdate(params models.UpdateVoucherParams) error {
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

// ValidateList validates voucher listing parameters
func (v *voucherValidator) ValidateList(params models.ListVouchersParams) error {
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
