-- name: CreateProduct :one
INSERT INTO products (
    store_id,
    name,
    slug,
    description,
    type,
    status,
    price,
    compare_at_price,
    cost_price,
    sku,
    barcode,
    weight,
    weight_unit,
    is_taxable,
    is_featured,
    is_gift_card,
    requires_shipping,
    inventory_quantity,
    inventory_policy,
    inventory_tracking,
    seo_title,
    seo_description,
    metadata
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetProductBySlug :one
SELECT * FROM products
WHERE store_id = $1 AND slug = $2 AND deleted_at IS NULL;

-- name: ListProducts :many
SELECT * FROM products
WHERE store_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListProductsByStatus :many
SELECT * FROM products
WHERE store_id = $1 AND status = $2 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: ListFeaturedProducts :many
SELECT * FROM products
WHERE store_id = $1 AND is_featured = true AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE($2, name),
    slug = COALESCE($3, slug),
    description = COALESCE($4, description),
    type = COALESCE($5, type),
    status = COALESCE($6, status),
    price = COALESCE($7, price),
    compare_at_price = COALESCE($8, compare_at_price),
    cost_price = COALESCE($9, cost_price),
    sku = COALESCE($10, sku),
    barcode = COALESCE($11, barcode),
    weight = COALESCE($12, weight),
    weight_unit = COALESCE($13, weight_unit),
    is_taxable = COALESCE($14, is_taxable),
    is_featured = COALESCE($15, is_featured),
    is_gift_card = COALESCE($16, is_gift_card),
    requires_shipping = COALESCE($17, requires_shipping),
    inventory_quantity = COALESCE($18, inventory_quantity),
    inventory_policy = COALESCE($19, inventory_policy),
    inventory_tracking = COALESCE($20, inventory_tracking),
    seo_title = COALESCE($21, seo_title),
    seo_description = COALESCE($22, seo_description),
    metadata = COALESCE($23, metadata)
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: CountProducts :one
SELECT COUNT(*) FROM products
WHERE store_id = $1 AND deleted_at IS NULL;

-- name: CountProductsByStatus :one
SELECT COUNT(*) FROM products
WHERE store_id = $1 AND status = $2 AND deleted_at IS NULL;

-- name: UpdateProductInventory :one
UPDATE products
SET
    inventory_quantity = inventory_quantity + $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;
