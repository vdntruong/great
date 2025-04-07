-- name: CreateProductVariant :one
INSERT INTO product_variants (
    product_id,
    name,
    sku,
    barcode,
    price,
    compare_at_price,
    cost_price,
    weight,
    weight_unit,
    inventory_quantity,
    inventory_policy,
    inventory_tracking,
    option_values
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetProductVariant :one
SELECT * FROM product_variants
WHERE id = $1;

-- name: ListProductVariants :many
SELECT * FROM product_variants
WHERE product_id = $1
ORDER BY created_at DESC;

-- name: UpdateProductVariant :one
UPDATE product_variants
SET
    name = COALESCE($2, name),
    sku = COALESCE($3, sku),
    barcode = COALESCE($4, barcode),
    price = COALESCE($5, price),
    compare_at_price = COALESCE($6, compare_at_price),
    cost_price = COALESCE($7, cost_price),
    weight = COALESCE($8, weight),
    weight_unit = COALESCE($9, weight_unit),
    inventory_quantity = COALESCE($10, inventory_quantity),
    inventory_policy = COALESCE($11, inventory_policy),
    inventory_tracking = COALESCE($12, inventory_tracking),
    option_values = COALESCE($13, option_values)
WHERE id = $1
RETURNING *;

-- name: DeleteProductVariant :exec
DELETE FROM product_variants
WHERE id = $1;

-- name: DeleteProductVariants :exec
DELETE FROM product_variants
WHERE product_id = $1;

-- name: UpdateVariantInventory :one
UPDATE product_variants
SET inventory_quantity = inventory_quantity + $2
WHERE id = $1
RETURNING *;

-- name: GetVariantBySKU :one
SELECT * FROM product_variants
WHERE product_id = $1 AND sku = $2;

-- name: GetVariantByBarcode :one
SELECT * FROM product_variants
WHERE product_id = $1 AND barcode = $2;

-- name: CountProductVariants :one
SELECT COUNT(*) FROM product_variants
WHERE product_id = $1;
