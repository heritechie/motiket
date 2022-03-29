// Code generated by sqlc. DO NOT EDIT.
// source: order_ticket.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createOrderTicket = `-- name: CreateOrderTicket :one
INSERT INTO order_ticket (
  qty,
  ticket_id,
  customer_order_id
) VALUES (
  $1, $2, $3
) RETURNING qty, ticket_id, customer_order_id
`

type CreateOrderTicketParams struct {
	Qty             int32     `json:"qty"`
	TicketID        uuid.UUID `json:"ticket_id"`
	CustomerOrderID uuid.UUID `json:"customer_order_id"`
}

func (q *Queries) CreateOrderTicket(ctx context.Context, arg CreateOrderTicketParams) (OrderTicket, error) {
	row := q.db.QueryRowContext(ctx, createOrderTicket, arg.Qty, arg.TicketID, arg.CustomerOrderID)
	var i OrderTicket
	err := row.Scan(&i.Qty, &i.TicketID, &i.CustomerOrderID)
	return i, err
}

const listOrderTicket = `-- name: ListOrderTicket :many
SELECT qty, ticket_id, customer_order_id FROM order_ticket
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOrderTicketParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrderTicket(ctx context.Context, arg ListOrderTicketParams) ([]OrderTicket, error) {
	rows, err := q.db.QueryContext(ctx, listOrderTicket, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderTicket{}
	for rows.Next() {
		var i OrderTicket
		if err := rows.Scan(&i.Qty, &i.TicketID, &i.CustomerOrderID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
