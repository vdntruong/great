-- name: AddProductCategory :exec
INSERT INTO product_categories (product_id, category_id)
VALUES ($1, $2);

-- name: RemoveProductCategory :exec
DELETE FROM product_categories
WHERE product_id = $1 AND category_id = $2;

-- name: GetProductCategories :many
SELECT c.* FROM store_categories c
JOIN product_categories pc ON c.id = pc.category_id
WHERE pc.product_id = $1;

-- name: GetCategoryProducts :many
SELECT p.* FROM products p
JOIN product_categories pc ON p.id = pc.product_id
WHERE pc.category_id = $1 AND p.deleted_at IS NULL
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountCategoryProducts :one
SELECT COUNT(*) FROM products p
JOIN product_categories pc ON p.id = pc.product_id
WHERE pc.category_id = $1 AND p.deleted_at IS NULL;

-- name: RemoveAllProductCategories :exec
DELETE FROM product_categories
WHERE product_id = $1;

-- name: UpdateProductCategories :exec
WITH new_categories AS (
    SELECT unnest($2::uuid[]) AS category_id
)
DELETE FROM product_categories pc
WHERE pc.product_id = $1
    AND pc.category_id NOT IN (SELECT category_id FROM new_categories);

INSERT INTO product_categories (product_id, category_id)
SELECT $1, category_id
FROM new_categories
WHERE NOT EXISTS (
    SELECT 1 FROM product_categories pc
    WHERE pc.product_id = $1 AND pc.category_id = new_categories.category_id
);
