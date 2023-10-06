// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: cart_item.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createCartItem = `-- name: CreateCartItem :one

INSERT INTO cart_items(
    cart_id,item_id,quantity
)
values ($1,$2,$3)
RETURNING cart_id, item_id, quantity
`

type CreateCartItemParams struct {
	CartID   int32         `json:"cart_id"`
	ItemID   int32         `json:"item_id"`
	Quantity sql.NullInt32 `json:"quantity"`
}

func (q *Queries) CreateCartItem(ctx context.Context, arg CreateCartItemParams) (CartItem, error) {
	row := q.db.QueryRowContext(ctx, createCartItem, arg.CartID, arg.ItemID, arg.Quantity)
	var i CartItem
	err := row.Scan(&i.CartID, &i.ItemID, &i.Quantity)
	return i, err
}

const deleteCartItem = `-- name: DeleteCartItem :one
DELETE FROM cart_items WHERE cart_id = $1 AND item_id = $2
RETURNING cart_id, item_id, quantity
`

type DeleteCartItemParams struct {
	CartID int32 `json:"cart_id"`
	ItemID int32 `json:"item_id"`
}

func (q *Queries) DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) (CartItem, error) {
	row := q.db.QueryRowContext(ctx, deleteCartItem, arg.CartID, arg.ItemID)
	var i CartItem
	err := row.Scan(&i.CartID, &i.ItemID, &i.Quantity)
	return i, err
}

const getCartItemFromCartID = `-- name: GetCartItemFromCartID :many
SELECT cart_id, item_id, quantity FROM cart_items
WHERE cart_id = $1
`

func (q *Queries) GetCartItemFromCartID(ctx context.Context, cartID int32) ([]CartItem, error) {
	rows, err := q.db.QueryContext(ctx, getCartItemFromCartID, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CartItem{}
	for rows.Next() {
		var i CartItem
		if err := rows.Scan(&i.CartID, &i.ItemID, &i.Quantity); err != nil {
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

const getCartItemFromCartIDAndItemID = `-- name: GetCartItemFromCartIDAndItemID :one
SELECT cart_id, item_id, quantity FROM cart_items
WHERE cart_id = $1 AND item_id = $2
`

type GetCartItemFromCartIDAndItemIDParams struct {
	CartID int32 `json:"cart_id"`
	ItemID int32 `json:"item_id"`
}

func (q *Queries) GetCartItemFromCartIDAndItemID(ctx context.Context, arg GetCartItemFromCartIDAndItemIDParams) (CartItem, error) {
	row := q.db.QueryRowContext(ctx, getCartItemFromCartIDAndItemID, arg.CartID, arg.ItemID)
	var i CartItem
	err := row.Scan(&i.CartID, &i.ItemID, &i.Quantity)
	return i, err
}

const updateCartItem = `-- name: UpdateCartItem :one
UPDATE cart_items
SET
    quantity = $1
WHERE cart_id = $3 AND item_id = $2
RETURNING cart_id, item_id, quantity
`

type UpdateCartItemParams struct {
	Quantity sql.NullInt32 `json:"quantity"`
	ItemID   int32         `json:"item_id"`
	CartID   int32         `json:"cart_id"`
}

func (q *Queries) UpdateCartItem(ctx context.Context, arg UpdateCartItemParams) (CartItem, error) {
	row := q.db.QueryRowContext(ctx, updateCartItem, arg.Quantity, arg.ItemID, arg.CartID)
	var i CartItem
	err := row.Scan(&i.CartID, &i.ItemID, &i.Quantity)
	return i, err
}
