-- name: CreateCart :one
INSERT INTO cart(
    username
)
values ($1)
RETURNING *;

-- name: UpdateCartAmount :one
UPDATE cart
SET
    total_value = $1
WHERE username = $2
RETURNING *;

-- name: GetCart :one
SELECT * FROM cart 
WHERE username = $1 LIMIT 1;

