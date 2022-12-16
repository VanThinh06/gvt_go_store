

-- name: CreateUser :one
INSERT INTO "user" (name,
                   address,
                   phone,
                   email,
                   type_user,
                   password,
                   payment,
                   payment_info,
                   payment_number)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9) RETURNING *;

-- name: CreateAccount :one
INSERT INTO "user" (
                   email,
                   type_user,
                   password)
VALUES ($1,
        $2,
        $3) RETURNING *;

-- name: ListUser :many

SELECT *
FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetUserById :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one

UPDATE "user"
set name = $2,
    payment_number = $3, payment_info = $4, address = $5, phone = $6, email = $7, type_user = $8, password = $9, payment = $10
WHERE id = $1 RETURNING *;

-- name: DeleteUser :one

DELETE
FROM "user"
WHERE id = $1 RETURNING *;

