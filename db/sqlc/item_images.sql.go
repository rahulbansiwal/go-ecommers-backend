// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: item_images.sql

package sqlc

import (
	"context"
)

const createItemImage = `-- name: CreateItemImage :one
INSERT INTO item_images(
    item_id,image_url
)
VALUES ($1,$2)
RETURNING id, item_id, image_url
`

type CreateItemImageParams struct {
	ItemID   int32  `json:"item_id"`
	ImageUrl string `json:"image_url"`
}

func (q *Queries) CreateItemImage(ctx context.Context, arg CreateItemImageParams) (ItemImage, error) {
	row := q.db.QueryRowContext(ctx, createItemImage, arg.ItemID, arg.ImageUrl)
	var i ItemImage
	err := row.Scan(&i.ID, &i.ItemID, &i.ImageUrl)
	return i, err
}

const deleteItemImage = `-- name: DeleteItemImage :one
DELETE FROM item_images WHERE id = $1 
RETURNING id, item_id, image_url
`

func (q *Queries) DeleteItemImage(ctx context.Context, id int32) (ItemImage, error) {
	row := q.db.QueryRowContext(ctx, deleteItemImage, id)
	var i ItemImage
	err := row.Scan(&i.ID, &i.ItemID, &i.ImageUrl)
	return i, err
}

const getItemImagesFromItemId = `-- name: GetItemImagesFromItemId :many
SELECT id, item_id, image_url FROM item_images
WHERE item_id = $1
`

func (q *Queries) GetItemImagesFromItemId(ctx context.Context, itemID int32) ([]ItemImage, error) {
	rows, err := q.db.QueryContext(ctx, getItemImagesFromItemId, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ItemImage{}
	for rows.Next() {
		var i ItemImage
		if err := rows.Scan(&i.ID, &i.ItemID, &i.ImageUrl); err != nil {
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

const updateItemImageURL = `-- name: UpdateItemImageURL :one
UPDATE item_images
SET
    image_url = $1
WHERE id = $2
RETURNING id, item_id, image_url
`

type UpdateItemImageURLParams struct {
	ImageUrl string `json:"image_url"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateItemImageURL(ctx context.Context, arg UpdateItemImageURLParams) (ItemImage, error) {
	row := q.db.QueryRowContext(ctx, updateItemImageURL, arg.ImageUrl, arg.ID)
	var i ItemImage
	err := row.Scan(&i.ID, &i.ItemID, &i.ImageUrl)
	return i, err
}
