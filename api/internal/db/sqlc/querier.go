// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	GetEvent(ctx context.Context, id uuid.UUID) (Event, error)
	ListEvent(ctx context.Context, arg ListEventParams) ([]Event, error)
}

var _ Querier = (*Queries)(nil)
