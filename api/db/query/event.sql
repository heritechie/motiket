-- name: CreateEvent :one
INSERT INTO event (
  id,
  name,
  description,
  start_date,
  end_date,
  prefix
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetEvent :one
SELECT * FROM event
WHERE id = $1 LIMIT 1;

-- name: ListEvent :many
SELECT * FROM event
ORDER BY id
LIMIT $1
OFFSET $2;