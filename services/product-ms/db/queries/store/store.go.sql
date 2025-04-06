-- name: GetStoreByID :one
SELECT * FROM stores WHERE id = $1 AND deleted_at IS NULL;

-- name: GetStoreBySlug :one
SELECT * FROM stores WHERE slug = $1 AND deleted_at IS NULL;

-- name: ListStores :many
SELECT * FROM stores WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: CreateStore :one
INSERT INTO stores (
    id, name, slug, description, logo_url, cover_url,
    status, is_verified, owner_id, contact_email, contact_phone,
    address, settings
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
    $11, $12, $13
) RETURNING *;

-- name: UpdateStore :one
UPDATE stores SET
    name = COALESCE($2, name),
    slug = COALESCE($3, slug),
    description = COALESCE($4, description),
    logo_url = COALESCE($5, logo_url),
    cover_url = COALESCE($6, cover_url),
    status = COALESCE($7, status),
    is_verified = COALESCE($8, is_verified),
    contact_email = COALESCE($9, contact_email),
    contact_phone = COALESCE($10, contact_phone),
    address = COALESCE($11, address),
    settings = COALESCE($12, settings),
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteStore :exec
UPDATE stores SET deleted_at = NOW() WHERE id = $1;

-- name: GetStoreCategories :many
SELECT * FROM store_categories WHERE store_id = $1 ORDER BY sort_order ASC;

-- name: GetStoreCategoryByID :one
SELECT * FROM store_categories WHERE id = $1;

-- name: GetStoreCategoryBySlug :one
SELECT * FROM store_categories WHERE store_id = $1 AND slug = $2;

-- name: CreateStoreCategory :one
INSERT INTO store_categories (
    id, store_id, name, slug, description,
    parent_id, sort_order
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: UpdateStoreCategory :one
UPDATE store_categories SET
    name = COALESCE($2, name),
    slug = COALESCE($3, slug),
    description = COALESCE($4, description),
    parent_id = COALESCE($5, parent_id),
    sort_order = COALESCE($6, sort_order),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteStoreCategory :exec
DELETE FROM store_categories WHERE id = $1;
