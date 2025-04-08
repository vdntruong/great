-- name: CreateStoreCategory :one
INSERT INTO store_categories (
    store_id,
    name,
    slug,
    description,
    parent_id,
    sort_order
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetStoreCategory :one
SELECT * FROM store_categories
WHERE id = $1;

-- name: GetStoreCategoryBySlug :one
SELECT * FROM store_categories
WHERE store_id = $1 AND slug = $2;

-- name: ListStoreCategories :many
SELECT * FROM store_categories
WHERE store_id = $1
ORDER BY sort_order ASC, created_at DESC;

-- name: ListStoreCategoriesByParent :many
SELECT * FROM store_categories
WHERE store_id = $1 AND parent_id = $2
ORDER BY sort_order ASC, created_at DESC;

-- name: ListRootStoreCategories :many
SELECT * FROM store_categories
WHERE store_id = $1 AND parent_id IS NULL
ORDER BY sort_order ASC, created_at DESC;

-- name: UpdateStoreCategory :one
UPDATE store_categories
SET
    name = COALESCE($2, name),
    slug = COALESCE($3, slug),
    description = COALESCE($4, description),
    parent_id = COALESCE($5, parent_id),
    sort_order = COALESCE($6, sort_order)
WHERE id = $1
RETURNING *;

-- name: DeleteStoreCategory :exec
DELETE FROM store_categories
WHERE id = $1;

-- name: CountStoreCategories :one
SELECT COUNT(*) FROM store_categories
WHERE store_id = $1;

-- name: CountStoreCategoriesByParent :one
SELECT COUNT(*) FROM store_categories
WHERE store_id = $1 AND parent_id = $2;
