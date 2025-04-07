-- name: CreateProductImage :one
INSERT INTO product_images (
    product_id,
    variant_id,
    url,
    alt_text,
    type,
    sort_order
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetProductImage :one
SELECT * FROM product_images
WHERE id = $1;

-- name: ListProductImages :many
SELECT * FROM product_images
WHERE product_id = $1
ORDER BY sort_order ASC, created_at DESC;

-- name: ListVariantImages :many
SELECT * FROM product_images
WHERE variant_id = $1
ORDER BY sort_order ASC, created_at DESC;

-- name: UpdateProductImage :one
UPDATE product_images
SET
    url = COALESCE($2, url),
    alt_text = COALESCE($3, alt_text),
    type = COALESCE($4, type),
    sort_order = COALESCE($5, sort_order)
WHERE id = $1
RETURNING *;

-- name: DeleteProductImage :exec
DELETE FROM product_images
WHERE id = $1;

-- name: DeleteProductImages :exec
DELETE FROM product_images
WHERE product_id = $1;

-- name: DeleteVariantImages :exec
DELETE FROM product_images
WHERE variant_id = $1;

-- name: UpdateImageSortOrder :exec
UPDATE product_images
SET sort_order = $2
WHERE id = $1;

-- name: GetMainProductImage :one
SELECT * FROM product_images
WHERE product_id = $1 AND type = 'main'
LIMIT 1;

-- name: GetThumbnailProductImage :one
SELECT * FROM product_images
WHERE product_id = $1 AND type = 'thumbnail'
LIMIT 1;

-- name: GetGalleryProductImages :many
SELECT * FROM product_images
WHERE product_id = $1 AND type = 'gallery'
ORDER BY sort_order ASC;
