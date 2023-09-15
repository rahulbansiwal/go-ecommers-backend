-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(
    username,hashed_password,full_name,mobile_number
)
values($1,$2,$3,$4)
RETURNING *;
