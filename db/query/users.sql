-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(
    username,hashed_password,full_name,mobile_number
)
values($1,$2,$3,$4)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    username = coalesce(sqlc.narg('username'),username),
    hashed_password = coalesce(sqlc.narg('hashed_password'),hashed_password),
    full_name =coalesce(sqlc.narg('full_name'),full_name),
    mobile_number = coalesce(sqlc.narg('mobile_number'),mobile_number)
WHERE username = sqlc.arg('user')
RETURNING *;

