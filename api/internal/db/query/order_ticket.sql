-- name: CreateOrderTicket :one
INSERT INTO order_ticket (
  qty,
  ticket_id,
  customer_order_id
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: ListOrderTicketByCustomerOrderId :many
SELECT 
  t.id, 
  t.serial_number, 
  COALESCE(t.purchase_date, now()) purchase_date,
  tc.name category_name
FROM order_ticket ot
INNER JOIN ticket t ON t.id = ot.ticket_id
INNER JOIN ticket_category tc ON tc.id = t.ticket_category_id
WHERE ot.customer_order_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;