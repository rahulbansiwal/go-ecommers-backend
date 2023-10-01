// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: item.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items(
    name,price,category,created_by
)
values ($1,$2,$3,$4)
RETURNING id, name, price, created_by, discount, category, created_at
`

type CreateItemParams struct {
	Name      string `json:"name"`
	Price     string `json:"price"`
	Category  string `json:"category"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Name,
		arg.Price,
		arg.Category,
		arg.CreatedBy,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.CreatedBy,
		&i.Discount,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :one
DELETE FROM items WHERE id = $1 
RETURNING id, name, price, created_by, discount, category, created_at
`

func (q *Queries) DeleteItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, deleteItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.CreatedBy,
		&i.Discount,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const getItemById = `-- name: GetItemById :one
SELECT id, name, price, created_by, discount, category, created_at FROM items WHERE id = $1 LIMIT 1
`

func (q *Queries) GetItemById(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemById, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.CreatedBy,
		&i.Discount,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const getItemByName = `-- name: GetItemByName :one
SELECT id, name, price, created_by, discount, category, created_at FROM items WHERE name = $1 LIMIT 1
`

func (q *Queries) GetItemByName(ctx context.Context, name string) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemByName, name)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.CreatedBy,
		&i.Discount,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET
    name = coalesce($1,name),
    price = coalesce($2,price),
    discount = coalesce($3,discount),
    category = coalesce($4,category)
WHERE id = $5
RETURNING id, name, price, created_by, discount, category, created_at
`

type UpdateItemParams struct {
	Name     sql.NullString `json:"name"`
	Price    sql.NullString `json:"price"`
	Discount sql.NullInt32  `json:"discount"`
	Category sql.NullString `json:"category"`
	ID       int32          `json:"id"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.Name,
		arg.Price,
		arg.Discount,
		arg.Category,
		arg.ID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.CreatedBy,
		&i.Discount,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}
