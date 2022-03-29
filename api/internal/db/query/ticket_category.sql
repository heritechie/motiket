-- name: CreateTicketCategory :one
INSERT INTO ticket_category (
  id,
  name,
  qty,
  price,
  start_date,
  end_date,
  prefix,
  area,
  event_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetTicketCategory :one
SELECT * FROM ticket_category
WHERE id = $1 LIMIT 1;

-- name: ListTicketCategory :many
SELECT * FROM ticket_category
ORDER BY id
LIMIT $1
OFFSET $2;