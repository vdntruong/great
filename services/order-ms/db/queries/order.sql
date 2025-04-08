-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    status,
    total_amount
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: CreateOrderItem :one
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
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrderItems :many
SELECT * FROM order_items WHERE order_id = $1;

-- name: ListOrders :many
SELECT o.* FROM orders o
WHERE
    ($1::UUID IS NULL OR o.user_id = $1) AND
    ($2::VARCHAR IS NULL OR o.status = $2)
ORDER BY o.created_at DESC
LIMIT $3 OFFSET $4;

-- name: UpdateOrderStatus :one
UPDATE orders
SET 
    status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
