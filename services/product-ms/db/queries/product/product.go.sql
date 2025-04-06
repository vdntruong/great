-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL;

-- name: GetProductBySlug :one
SELECT * FROM products WHERE store_id = $1 AND slug = $2 AND deleted_at IS NULL;

-- name: ListProducts :many
SELECT * FROM products WHERE store_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: ListProductsByCategory :many
SELECT p.* FROM products p
JOIN product_categories pc ON p.id = pc.product_id
WHERE pc.category_id = $1 AND p.deleted_at IS NULL
ORDER BY p.created_at DESC LIMIT $2 OFFSET $3;

-- name: CreateProduct :one
INSERT INTO products (
    id, store_id, name, slug, description, type,
    status, price, compare_at_price, cost_price, sku, barcode,
    weight, weight_unit, is_taxable, is_featured, is_gift_card,
    requires_shipping, inventory_quantity, inventory_policy,
    inventory_tracking, seo_title, seo_description, metadata
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,
    $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
) RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
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
    metadata = COALESCE($23, metadata),
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteProduct :exec
UPDATE products SET deleted_at = NOW() WHERE id = $1;

-- name: GetProductVariants :many
SELECT * FROM product_variants WHERE product_id = $1 ORDER BY created_at ASC;

-- name: GetProductVariantByID :one
SELECT * FROM product_variants WHERE id = $1;

-- name: CreateProductVariant :one
INSERT INTO product_variants (
    id, product_id, name, sku, price, compare_at_price,
    weight, weight_unit, inventory_quantity, option_values
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: UpdateProductVariant :one
UPDATE product_variants SET
    name = COALESCE($2, name),
    sku = COALESCE($3, sku),
    price = COALESCE($4, price),
    compare_at_price = COALESCE($5, compare_at_price),
    weight = COALESCE($6, weight),
    weight_unit = COALESCE($7, weight_unit),
    inventory_quantity = COALESCE($8, inventory_quantity),
    option_values = COALESCE($9, option_values),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProductVariant :exec
DELETE FROM product_variants WHERE id = $1;

-- name: GetProductImages :many
SELECT * FROM product_images WHERE product_id = $1 ORDER BY sort_order ASC;

-- name: GetProductImageByID :one
SELECT * FROM product_images WHERE id = $1;

-- name: CreateProductImage :one
INSERT INTO product_images (
    id, product_id, variant_id, url, alt_text,
    type, sort_order
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: UpdateProductImage :one
UPDATE product_images SET
    url = COALESCE($2, url),
    alt_text = COALESCE($3, alt_text),
    type = COALESCE($4, type),
    sort_order = COALESCE($5, sort_order),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProductImage :exec
DELETE FROM product_images WHERE id = $1;
