-- name: CreateTransaction :one

INSERT INTO "transaction" (id_user,
                         status,
                         amount,
                         message)
VALUES ($1,
        $2,
        $3,
        $4) RETURNING *;

-- name: ListTransaction :many

SELECT *
FROM "transaction"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetTransactionById :one
SELECT * FROM "transaction"
WHERE id = $1 LIMIT 1;

-- name: UpdateTransaction :one

UPDATE "transaction"
set id_user = $2, status = $3, amount = $4, message = $5
WHERE id = $1 RETURNING *;

-- name: DeleteTransaction :one

DELETE
FROM "transaction"
WHERE id = $1 RETURNING *;