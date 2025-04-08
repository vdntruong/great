package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"order-ms/db/dao"
	"order-ms/internal/models"
	"order-ms/internal/service/validator"

	"github.com/google/uuid"
)

type OrderValidator interface {
	ValidateCreate(params models.CreateOrderParams) error
	ValidateUpdate(params models.UpdateOrderStatusParams) error
	ValidateList(params models.ListOrdersParams) error
}

type OrderDAOQuerier interface {
	CreateOrder(ctx context.Context, arg dao.CreateOrderParams) (dao.Order, error)
	CreateOrderItem(ctx context.Context, arg dao.CreateOrderItemParams) (dao.OrderItem, error)
	GetOrder(ctx context.Context, id uuid.UUID) (dao.Order, error)
	GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]dao.OrderItem, error)
	ListOrders(ctx context.Context, arg dao.ListOrdersParams) ([]dao.Order, error)
	UpdateOrderStatus(ctx context.Context, arg dao.UpdateOrderStatusParams) (dao.Order, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}

type OrderServiceAdapter struct {
	querier   OrderDAOQuerier
	validator OrderValidator
}

var _ OrderService = (*OrderServiceAdapter)(nil)

func NewOrderServiceAdapter(querier OrderDAOQuerier) *OrderServiceAdapter {
	return &OrderServiceAdapter{
		querier:   querier,
		validator: validator.NewOrderValidator(),
	}
}

func (s *OrderServiceAdapter) CreateOrder(ctx context.Context, params models.CreateOrderParams) (*models.Order, error) {
	// Validate params
	if err := s.validator.ValidateCreate(params); err != nil {
		return nil, err
	}

	// Create order
	order, err := s.querier.CreateOrder(ctx, dao.CreateOrderParams{
		UserID: params.UserID,
		Status: string(models.OrderStatusPending),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Create order items
	for _, item := range params.Items {
		_, err := s.querier.CreateOrderItem(ctx, dao.CreateOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
		if err != nil {
			return nil, err
		}
	}

	// Get order items
	items, err := s.querier.GetOrderItems(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	// Convert to model
	result := &models.Order{
		ID:     order.ID,
		UserID: order.UserID,
		Status: models.OrderStatus(order.Status),
		Items:  make([]models.OrderItem, 0),
	}

	// Convert items
	for _, item := range items {
		result.Items = append(result.Items, models.OrderItem{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return result, nil
}

func (s *OrderServiceAdapter) GetOrder(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	// Get order
	order, err := s.querier.GetOrder(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	// Get order items
	items, err := s.querier.GetOrderItems(ctx, id)
	if err != nil {
		return nil, err
	}

	// Convert to model
	result := &models.Order{
		ID:     order.ID,
		UserID: order.UserID,
		Status: models.OrderStatus(order.Status),
		Items:  make([]models.OrderItem, 0),
	}

	// Convert items
	for _, item := range items {
		result.Items = append(result.Items, models.OrderItem{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return result, nil
}

func (s *OrderServiceAdapter) ListOrders(ctx context.Context, params models.ListOrdersParams) ([]*models.Order, error) {
	// Validate params
	if err := s.validator.ValidateList(params); err != nil {
		return nil, err
	}

	// List orders
	orders, err := s.querier.ListOrders(ctx, dao.ListOrdersParams{
		Limit:  int32(params.PageSize),
		Offset: int32((params.Page - 1) * params.PageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}

	// Convert to models
	result := make([]*models.Order, 0)
	for _, order := range orders {
		// Get order items
		items, err := s.querier.GetOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}

		// Convert to model
		orderModel := &models.Order{
			ID:     order.ID,
			UserID: order.UserID,
			Status: models.OrderStatus(order.Status),
			Items:  make([]models.OrderItem, 0),
		}

		// Convert items
		for _, item := range items {
			orderModel.Items = append(orderModel.Items, models.OrderItem{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
		}

		result = append(result, orderModel)
	}

	return result, nil
}

func (s *OrderServiceAdapter) UpdateOrderStatus(ctx context.Context, params models.UpdateOrderStatusParams) (*models.Order, error) {
	// Validate params
	if err := s.validator.ValidateUpdate(params); err != nil {
		return nil, err
	}

	// Update order
	order, err := s.querier.UpdateOrderStatus(ctx, dao.UpdateOrderStatusParams{
		ID:     params.ID,
		Status: string(params.Status),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %w", err)
	}

	// Get order items
	items, err := s.querier.GetOrderItems(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	// Convert to model
	result := &models.Order{
		ID:     order.ID,
		UserID: order.UserID,
		Status: models.OrderStatus(order.Status),
		Items:  make([]models.OrderItem, 0),
	}

	// Convert items
	for _, item := range items {
		result.Items = append(result.Items, models.OrderItem{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return result, nil
}

func (s *OrderServiceAdapter) CancelOrder(ctx context.Context, id uuid.UUID) error {
	// Validate order ID
	if err := s.validator.ValidateUpdate(models.UpdateOrderStatusParams{
		ID:     id,
		Status: models.OrderStatusCanceled,
	}); err != nil {
		return err
	}

	// Get order
	order, err := s.querier.GetOrder(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("order not found")
		}
		return err
	}

	// Check if order can be cancelled
	if order.Status != string(models.OrderStatusPending) {
		return errors.New("order cannot be cancelled")
	}

	// Update order status to cancelled
	_, err = s.querier.UpdateOrderStatus(ctx, dao.UpdateOrderStatusParams{
		ID:     id,
		Status: string(models.OrderStatusCanceled),
	})
	return err
}
