-- name: CreateOrderTicket :one
INSERT INTO order_ticket (
  qty,
  ticket_id,
  customer_order_id
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: ListOrderTicket :many
SELECT * FROM order_ticket
ORDER BY id
LIMIT $1
OFFSET $2;