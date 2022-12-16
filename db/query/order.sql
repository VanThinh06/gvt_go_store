-- name: CreateOrder :one

INSERT INTO "order" ( id_transaction,
        id_product,
        amount,
        data,
        status,
        qty)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING *;

-- name: ListOrder :many

SELECT *
FROM
"order"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetOrderById :one

SELECT *
FROM
"order"
WHERE id = $1
LIMIT 1;

-- name: UpdateOrder :one

UPDATE "order" set id_transaction = $2,
    id_product = $3,
    amount = $4,
    data =$5,
          status=$6,
          qty=$7
WHERE id = $1 RETURNING *;

-- name: DeleteOrder :one

DELETE
FROM "order"
WHERE id = $1 RETURNING *;