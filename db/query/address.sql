-- name: AddAddress :one
INSERT INTO addresses(
    username,full_name,country_code,city,street,landmark,mobile_number
)
values($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: UpdateAddress :one
UPDATE addresses
SET
    full_name = coalesce(sqlc.narg('full_name'),full_name),
    country_code = coalesce(sqlc.narg('country_code'),country_code),
    city = coalesce(sqlc.narg('city'),city),
    street = coalesce(sqlc.narg('street'),street),
    landmark = coalesce(sqlc.narg('landmark'),landmark),
    mobile_number = coalesce(sqlc.narg('mobile_number'),mobile_number)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteAddress :one
DELETE FROM addresses WHERE id = $1 RETURNING *;