-- name: CreateStore :one
INSERT INTO stores (
    name,
    slug,
    description,
    logo_url,
    cover_url,
    status,
    is_verified,
    owner_id,
    contact_email,
    contact_phone,
    address,
    settings
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetStoreBySlug :one
SELECT * FROM stores
WHERE slug = $1 AND deleted_at IS NULL;

-- name: ListStores :many
SELECT * FROM stores
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListStoresByOwner :many
SELECT * FROM stores
WHERE owner_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateStore :one
UPDATE stores
SET
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
    settings = COALESCE($12, settings)
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteStore :exec
UPDATE stores
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: CountStores :one
SELECT COUNT(*) FROM stores
WHERE deleted_at IS NULL;

-- name: CountStoresByOwner :one
SELECT COUNT(*) FROM stores
WHERE owner_id = $1 AND deleted_at IS NULL;
