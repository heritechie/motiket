-- name: CreateCustomer :one
INSERT INTO customer (
  id,
  full_name,
  email,
  password,
  phone_number,
  confirmation_code,
  confirmation_time
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customer
WHERE id = $1 LIMIT 1;

-- name: ListCustomer :many
SELECT * FROM customer
ORDER BY id
LIMIT $1
OFFSET $2;