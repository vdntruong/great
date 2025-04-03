-- name: CreateProduct :one
INSERT INTO products (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
