-- name: CreateVoucher :one
INSERT INTO vouchers (
    store_id,
    code,
    type,
    value,
    min_purchase_amount,
    max_discount_amount,
    start_date,
    end_date,
    usage_limit,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetVoucher :one
SELECT * FROM vouchers
WHERE id = $1;

-- name: GetVoucherByCode :one
SELECT * FROM vouchers
WHERE store_id = $1 AND code = $2;

-- name: ListVouchers :many
SELECT * FROM vouchers
WHERE store_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListActiveVouchers :many
SELECT * FROM vouchers
WHERE store_id = $1
    AND status = 'active'
    AND (end_date IS NULL OR end_date > CURRENT_TIMESTAMP)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateVoucher :one
UPDATE vouchers
SET
    code = COALESCE($2, code),
    type = COALESCE($3, type),
    value = COALESCE($4, value),
    min_purchase_amount = COALESCE($5, min_purchase_amount),
    max_discount_amount = COALESCE($6, max_discount_amount),
    start_date = COALESCE($7, start_date),
    end_date = COALESCE($8, end_date),
    usage_limit = COALESCE($9, usage_limit),
    status = COALESCE($10, status)
WHERE id = $1
RETURNING *;

-- name: DeleteVoucher :exec
DELETE FROM vouchers
WHERE id = $1;

-- name: CountVouchers :one
SELECT COUNT(*) FROM vouchers
WHERE store_id = $1;

-- name: CountActiveVouchers :one
SELECT COUNT(*) FROM vouchers
WHERE store_id = $1
    AND status = 'active'
    AND (end_date IS NULL OR end_date > CURRENT_TIMESTAMP);

-- name: IncrementVoucherUsage :one
UPDATE vouchers
SET usage_count = usage_count + 1
WHERE id = $1 AND (usage_limit IS NULL OR usage_count < usage_limit)
RETURNING *;

-- name: AddVoucherProduct :exec
INSERT INTO voucher_products (voucher_id, product_id)
VALUES ($1, $2);

-- name: RemoveVoucherProduct :exec
DELETE FROM voucher_products
WHERE voucher_id = $1 AND product_id = $2;

-- name: AddVoucherCategory :exec
INSERT INTO voucher_categories (voucher_id, category_id)
VALUES ($1, $2);

-- name: RemoveVoucherCategory :exec
DELETE FROM voucher_categories
WHERE voucher_id = $1 AND category_id = $2;

-- name: GetVoucherProducts :many
SELECT p.* FROM products p
JOIN voucher_products vp ON p.id = vp.product_id
WHERE vp.voucher_id = $1;

-- name: GetVoucherCategories :many
SELECT c.* FROM store_categories c
JOIN voucher_categories vc ON c.id = vc.category_id
WHERE vc.voucher_id = $1;
