// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order.sql

package dao

import (
	"context"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    status,
    total_amount
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, user_id, status, total_amount, created_at, updated_at
`

type CreateOrderParams struct {
	UserID      uuid.UUID `json:"user_id"`
	Status      string    `json:"status"`
	TotalAmount string    `json:"total_amount"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.queryRow(ctx, q.createOrderStmt, createOrder, arg.UserID, arg.Status, arg.TotalAmount)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price,
    subtotal
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id, order_id, product_id, quantity, price, subtotal, created_at, updated_at
`

type CreateOrderItemParams struct {
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	Price     string    `json:"price"`
	Subtotal  string    `json:"subtotal"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.queryRow(ctx, q.createOrderItemStmt, createOrderItem,
		arg.OrderID,
		arg.ProductID,
		arg.Quantity,
		arg.Price,
		arg.Subtotal,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
		&i.Subtotal,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteOrderStmt, deleteOrder, id)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders WHERE id = $1
`

func (q *Queries) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.queryRow(ctx, q.getOrderStmt, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderItems = `-- name: GetOrderItems :many
SELECT id, order_id, product_id, quantity, price, subtotal, created_at, updated_at FROM order_items WHERE order_id = $1
`

func (q *Queries) GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]OrderItem, error) {
	rows, err := q.query(ctx, q.getOrderItemsStmt, getOrderItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
			&i.Price,
			&i.Subtotal,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrders = `-- name: ListOrders :many
SELECT o.id, o.user_id, o.status, o.total_amount, o.created_at, o.updated_at FROM orders o
WHERE
    ($1::UUID IS NULL OR o.user_id = $1) AND
    ($2::VARCHAR IS NULL OR o.status = $2)
ORDER BY o.created_at DESC
LIMIT $3 OFFSET $4
`

type ListOrdersParams struct {
	Column1 uuid.UUID `json:"column_1"`
	Column2 string    `json:"column_2"`
	Limit   int32     `json:"limit"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.query(ctx, q.listOrdersStmt, listOrders,
		arg.Column1,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Status,
			&i.TotalAmount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderStatus = `-- name: UpdateOrderStatus :one
UPDATE orders
SET 
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, status, total_amount, created_at, updated_at
`

type UpdateOrderStatusParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (Order, error) {
	row := q.queryRow(ctx, q.updateOrderStatusStmt, updateOrderStatus, arg.ID, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
