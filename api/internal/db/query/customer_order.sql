-- name: CreateCustomerOrder :one
INSERT INTO customer_order (
  id,
  order_time,
  time_paid,
  total_price,
  discount,
  final_price,
  customer_id,
  customer_payment_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCustomerOrder :one
SELECT * FROM customer_order
WHERE id = $1 LIMIT 1;

-- name: ListCustomerOrder :many
SELECT * FROM customer_order
ORDER BY id
LIMIT $1
OFFSET $2;