-- name: CreateTicket :one
INSERT INTO ticket (
  id,
  serial_number,
  purchase_date,
  ticket_category_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTicket :one
SELECT * FROM ticket
WHERE id = $1 LIMIT 1;

-- name: ListTicket :many
SELECT * FROM ticket
ORDER BY id
LIMIT $1
OFFSET $2;