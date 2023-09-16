-- name: CreateItem :one
INSERT INTO items(
    name,price,created_at,discount,category,created_by
)
values ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetItemById :one
SELECT * FROM items WHERE id = $1 LIMIT 1;

-- name: DeleteItem :one
DELETE FROM items WHERE id = $1 
RETURNING *;

-- name: UpdateItem :one
UPDATE items
SET
    name = coalesce(sqlc.narg('name'),name),
    price = coalesce(sqlc.narg('price'),price),
    discount = coalesce(sqlc.narg('discount'),discount),
    category = coalesce(sqlc.narg('category'),category)
WHERE id = sqlc.arg('id')
RETURNING *;