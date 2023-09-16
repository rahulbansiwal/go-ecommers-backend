-- name: CreateItemImage :one
INSERT INTO item_images(
    item_id,image_url
)
VALUES ($1,$2)
RETURNING *;

-- name: DeleteItemImage :one
DELETE FROm item_images WHERE id = $1 
RETURNING *;

-- name: UpdateItemImageURL :one
UPDATE item_images
SET
    image_url = $1
WHERE id = $2
RETURNING *;
    