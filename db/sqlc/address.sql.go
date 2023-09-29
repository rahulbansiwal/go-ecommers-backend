// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: address.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addAddress = `-- name: AddAddress :one
INSERT INTO addresses(
    username,full_name,country_code,city,street,landmark,mobile_number
)
values($1,$2,$3,$4,$5,$6,$7)
RETURNING id, username, full_name, country_code, city, street, landmark, mobile_number
`

type AddAddressParams struct {
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	CountryCode  string `json:"country_code"`
	City         string `json:"city"`
	Street       string `json:"street"`
	Landmark     string `json:"landmark"`
	MobileNumber int64  `json:"mobile_number"`
}

func (q *Queries) AddAddress(ctx context.Context, arg AddAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, addAddress,
		arg.Username,
		arg.FullName,
		arg.CountryCode,
		arg.City,
		arg.Street,
		arg.Landmark,
		arg.MobileNumber,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.CountryCode,
		&i.City,
		&i.Street,
		&i.Landmark,
		&i.MobileNumber,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :one
DELETE FROM addresses WHERE id = $1 RETURNING id, username, full_name, country_code, city, street, landmark, mobile_number
`

func (q *Queries) DeleteAddress(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRowContext(ctx, deleteAddress, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.CountryCode,
		&i.City,
		&i.Street,
		&i.Landmark,
		&i.MobileNumber,
	)
	return i, err
}

const getAddresses = `-- name: GetAddresses :many
SELECT id, username, full_name, country_code, city, street, landmark, mobile_number FROM addresses WHERE username = $1
`

func (q *Queries) GetAddresses(ctx context.Context, username string) ([]Address, error) {
	rows, err := q.db.QueryContext(ctx, getAddresses, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Address{}
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FullName,
			&i.CountryCode,
			&i.City,
			&i.Street,
			&i.Landmark,
			&i.MobileNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getaddressById = `-- name: GetaddressById :one
SELECT id, username, full_name, country_code, city, street, landmark, mobile_number FROM addresses WHERE id = $1 LIMIT 1
`

func (q *Queries) GetaddressById(ctx context.Context, id int32) (Address, error) {
	row := q.db.QueryRowContext(ctx, getaddressById, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.CountryCode,
		&i.City,
		&i.Street,
		&i.Landmark,
		&i.MobileNumber,
	)
	return i, err
}

const updateAddress = `-- name: UpdateAddress :one
UPDATE addresses
SET
    full_name = coalesce($1,full_name),
    country_code = coalesce($2,country_code),
    city = coalesce($3,city),
    street = coalesce($4,street),
    landmark = coalesce($5,landmark),
    mobile_number = coalesce($6,mobile_number)
WHERE id = $7
RETURNING id, username, full_name, country_code, city, street, landmark, mobile_number
`

type UpdateAddressParams struct {
	FullName     sql.NullString `json:"full_name"`
	CountryCode  sql.NullString `json:"country_code"`
	City         sql.NullString `json:"city"`
	Street       sql.NullString `json:"street"`
	Landmark     sql.NullString `json:"landmark"`
	MobileNumber sql.NullInt64  `json:"mobile_number"`
	ID           int32          `json:"id"`
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, updateAddress,
		arg.FullName,
		arg.CountryCode,
		arg.City,
		arg.Street,
		arg.Landmark,
		arg.MobileNumber,
		arg.ID,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.CountryCode,
		&i.City,
		&i.Street,
		&i.Landmark,
		&i.MobileNumber,
	)
	return i, err
}
