-- name: CreateCustomerPayment :one
INSERT INTO customer_payment (
  id,
  status,
  success_at,
  failed_reason,
  customer_id,
  payment_option_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetCustomerPayment :one
SELECT * FROM customer_payment
WHERE id = $1 LIMIT 1;

-- name: ListCustomerPayment :many
SELECT * FROM customer_payment
ORDER BY id
LIMIT $1
OFFSET $2;