-- name: CreateSession :one
INSERT INTO sessions(
    id,username,refresh_token,client_ip,expired_at,is_blocked
)
values ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetSessionFromId :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: UpdateSession :one
UPDATE sessions
SET 
    is_blocked = $1
WHERE   id = $2
RETURNING *;

-- name: DeleteSessionById :one
DELETE FROM sessions WHERE id = $1
RETURNING *;

-- name: DeleteSessionByUsername :many
DELETE FROM sessions WHERE username = $1
RETURNING *;