-- name: CreateCategory :one

INSERT INTO category ( name, national)
VALUES ( $1,
         $2) RETURNING *;

-- name: ListCategory :many
SELECT *
FROM category
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetAllCategory :many
SELECT *
FROM category
ORDER BY id;

-- name: GetCategoryById :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: UpdateCategory :one

UPDATE category
set name = $2,
    national = $3
WHERE id = $1 RETURNING *;

-- name: DeleteCategory :one

DELETE
FROM category
WHERE id = $1 RETURNING *;