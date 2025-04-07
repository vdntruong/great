-- name: GetDiscountByID :one
SELECT * FROM discounts WHERE id = $1;

-- name: GetDiscountByCode :one
SELECT * FROM discounts WHERE store_id = $1 AND code = $2;

-- name: ListDiscounts :many
SELECT * FROM discounts WHERE store_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountDiscounts :one
SELECT COUNT(*) FROM discounts WHERE store_id = $1;

-- name: CreateDiscount :one
INSERT INTO discounts (
    id, store_id, name, code, type, value,
    scope, start_date, end_date, min_purchase_amount,
    max_discount_amount, usage_limit, is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
    $11, $12, $13
) RETURNING *;

-- name: UpdateDiscount :one
UPDATE discounts SET
    name = COALESCE(NULLIF($2, ''), name),
    code = COALESCE(NULLIF($3, ''), code),
    type = COALESCE(NULLIF($4, '')::discount_type, type),
    value = COALESCE(NULLIF($5, ''), value),
    scope = COALESCE(NULLIF($6, '')::discount_scope, scope),
    start_date = COALESCE($7, start_date),
    end_date = COALESCE($8, end_date),
    min_purchase_amount = COALESCE(NULLIF($9, ''), min_purchase_amount),
    max_discount_amount = COALESCE(NULLIF($10, ''), max_discount_amount),
    usage_limit = COALESCE($11, usage_limit),
    is_active = COALESCE($12, is_active),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDiscount :exec
DELETE FROM discounts WHERE id = $1;

-- name: GetDiscountProducts :many
SELECT p.* FROM products p
JOIN discount_products dp ON p.id = dp.product_id
WHERE dp.discount_id = $1;

-- name: AddDiscountProduct :exec
INSERT INTO discount_products (discount_id, product_id) VALUES ($1, $2);

-- name: RemoveDiscountProduct :exec
DELETE FROM discount_products WHERE discount_id = $1 AND product_id = $2;

-- name: GetDiscountCategories :many
SELECT c.* FROM store_categories c
JOIN discount_categories dc ON c.id = dc.category_id
WHERE dc.discount_id = $1;

-- name: AddDiscountCategory :exec
INSERT INTO discount_categories (discount_id, category_id) VALUES ($1, $2);

-- name: RemoveDiscountCategory :exec
DELETE FROM discount_categories WHERE discount_id = $1 AND category_id = $2;
