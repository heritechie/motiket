-- name: CreatePaymentOption :one
INSERT INTO payment_option (
  name
) VALUES ($1) RETURNING *;

-- name: GetPaymentOption :one
SELECT * FROM payment_option
WHERE id = $1 LIMIT 1;

-- name: ListPaymentOption :many
SELECT * FROM payment_option
ORDER BY id
LIMIT $1
OFFSET $2;