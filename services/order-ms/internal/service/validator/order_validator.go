package validator

import (
	"fmt"
	"order-ms/internal/models"

	"github.com/google/uuid"
)

type OrderValidator struct{}

func NewOrderValidator() *OrderValidator {
	return &OrderValidator{}
}

func (v *OrderValidator) ValidateCreate(params models.CreateOrderParams) error {
	if params.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}

	if len(params.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}

	for _, item := range params.Items {
		if item.ProductID == uuid.Nil {
			return fmt.Errorf("product_id is required for all items")
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("quantity must be greater than 0")
		}
	}

	return nil
}

func (v *OrderValidator) ValidateUpdate(params models.UpdateOrderStatusParams) error {
	if params.ID == uuid.Nil {
		return fmt.Errorf("order_id is required")
	}

	validStatuses := map[models.OrderStatus]bool{
		models.OrderStatusPending:   true,
		models.OrderStatusPaid:      true,
		models.OrderStatusShipping:  true,
		models.OrderStatusDelivered: true,
		models.OrderStatusCanceled:  true,
	}

	if !validStatuses[params.Status] {
		return fmt.Errorf("invalid status: %s", params.Status)
	}

	return nil
}

func (v *OrderValidator) ValidateList(params models.ListOrdersParams) error {
	if params.Page < 1 {
		return fmt.Errorf("page must be greater than 0")
	}

	if params.PageSize < 1 || params.PageSize > 100 {
		return fmt.Errorf("page_size must be between 1 and 100")
	}

	if params.Status != nil {
		validStatuses := map[models.OrderStatus]bool{
			models.OrderStatusPending:   true,
			models.OrderStatusPaid:      true,
			models.OrderStatusShipping:  true,
			models.OrderStatusDelivered: true,
			models.OrderStatusCanceled:  true,
		}

		if !validStatuses[*params.Status] {
			return fmt.Errorf("invalid status: %s", *params.Status)
		}
	}

	return nil
}
