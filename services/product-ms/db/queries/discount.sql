-- name: CreateDiscount :one
INSERT INTO discounts (
    store_id,
    name,
    code,
    type,
    value,
    scope,
    start_date,
    end_date,
    min_purchase_amount,
    max_discount_amount,
    usage_limit,
    is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetDiscount :one
SELECT * FROM discounts
WHERE id = $1;

-- name: GetDiscountByCode :one
SELECT * FROM discounts
WHERE store_id = $1 AND code = $2;

-- name: ListDiscounts :many
SELECT * FROM discounts
WHERE store_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListActiveDiscounts :many
SELECT * FROM discounts
WHERE store_id = $1
    AND is_active = true
    AND (end_date IS NULL OR end_date > CURRENT_TIMESTAMP)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateDiscount :one
UPDATE discounts
SET
    name = COALESCE($2, name),
    code = COALESCE($3, code),
    type = COALESCE($4, type),
    value = COALESCE($5, value),
    scope = COALESCE($6, scope),
    start_date = COALESCE($7, start_date),
    end_date = COALESCE($8, end_date),
    min_purchase_amount = COALESCE($9, min_purchase_amount),
    max_discount_amount = COALESCE($10, max_discount_amount),
    usage_limit = COALESCE($11, usage_limit),
    is_active = COALESCE($12, is_active)
WHERE id = $1
RETURNING *;

-- name: DeleteDiscount :exec
DELETE FROM discounts
WHERE id = $1;

-- name: CountDiscounts :one
SELECT COUNT(*) FROM discounts
WHERE store_id = $1;

-- name: CountActiveDiscounts :one
SELECT COUNT(*) FROM discounts
WHERE store_id = $1
    AND is_active = true
    AND (end_date IS NULL OR end_date > CURRENT_TIMESTAMP);

-- name: IncrementDiscountUsage :one
UPDATE discounts
SET usage_count = usage_count + 1
WHERE id = $1 AND (usage_limit IS NULL OR usage_count < usage_limit)
RETURNING *;

-- name: AddDiscountProduct :exec
INSERT INTO discount_products (discount_id, product_id)
VALUES ($1, $2);

-- name: RemoveDiscountProduct :exec
DELETE FROM discount_products
WHERE discount_id = $1 AND product_id = $2;

-- name: AddDiscountCategory :exec
INSERT INTO discount_categories (discount_id, category_id)
VALUES ($1, $2);

-- name: RemoveDiscountCategory :exec
DELETE FROM discount_categories
WHERE discount_id = $1 AND category_id = $2;

-- name: GetDiscountProducts :many
SELECT p.* FROM products p
JOIN discount_products dp ON p.id = dp.product_id
WHERE dp.discount_id = $1;

-- name: GetDiscountCategories :many
SELECT c.* FROM store_categories c
JOIN discount_categories dc ON c.id = dc.category_id
WHERE dc.discount_id = $1;
