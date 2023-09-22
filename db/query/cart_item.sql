-- name: CreateCartItem :one

INSERT INTO cart_items(
    cart_id,item_id
)
values ($1,$2)
RETURNING *;

-- name: UpdateCartItem :one
UPDATE cart_items
SET
    quantity = $1
WHERE cart_id = $1 AND item_id = $2
RETURNING *;

-- name: GetCartItemFromCartID :many
SELECT * FROM cart_items
WHERE cart_id = $1;

-- name: DeleteCartItem :one
DELETE FROM cart_items WHERE cart_id = $1 AND item_id = $2
RETURNING *;