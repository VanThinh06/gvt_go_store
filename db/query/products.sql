-- name: CreateProduct :one
INSERT INTO product (id_category, name, price, image, list_image, description, sold, status, sale)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9) RETURNING *;

-- name: ListProduct :many

SELECT *
FROM product
ORDER BY id
LIMIT $1
OFFSET $2;

-- -- name: GetProductWithCategory :many

-- SELECT product.id,  product.id_category, product.name, product.price, product.image, product.list_image, product.description, product.sold, product.status, product.sale, product.created_at, product.update_at
-- FROM  product
-- INNER JOIN category ON product.id_category = category.id
-- WHERE product.id = $1
-- ORDER BY product.id
-- LIMIT $2
-- OFFSET $3;


-- name: GetAllProduct :many

SELECT *
FROM product
ORDER BY id;

-- name: GetProductById :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: UpdateProduct :one

UPDATE product
set id_category = $2, name = $3, price = $4, image = $5, list_image = $6, description = $7, sold = $8, status= $9, sale=$10
WHERE id = $1 RETURNING *;

-- name: DeleteProduct :one

DELETE
FROM product
WHERE id = $1 RETURNING *;