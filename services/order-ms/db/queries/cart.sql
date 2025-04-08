-- name: CreateCart :one
INSERT INTO carts (
    user_id
) VALUES (
    $1
) RETURNING *;

-- name: GetCart :one
SELECT c.*, 
    COALESCE(
        json_agg(
            json_build_object(
                'id', ci.id,
                'product_id', ci.product_id,
                'quantity', ci.quantity
            )
        ) FILTER (WHERE ci.id IS NOT NULL),
        '[]'
    ) as items
FROM carts c
LEFT JOIN cart_items ci ON c.id = ci.cart_id
WHERE c.user_id = $1
GROUP BY c.id;

-- name: AddCartItem :one
INSERT INTO cart_items (
    cart_id,
    product_id,
    quantity
) VALUES (
    $1,
    $2,
    $3
) ON CONFLICT (cart_id, product_id) DO UPDATE
SET quantity = cart_items.quantity + EXCLUDED.quantity,
    updated_at = NOW()
RETURNING *;

-- name: UpdateCartItem :one
UPDATE cart_items
SET 
    quantity = $3,
    updated_at = NOW()
WHERE cart_id = $1 AND product_id = $2
RETURNING *;

-- name: RemoveCartItem :exec
DELETE FROM cart_items
WHERE cart_id = $1 AND product_id = $2;

-- name: ClearCart :exec
DELETE FROM cart_items WHERE cart_id = $1;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = $1;
