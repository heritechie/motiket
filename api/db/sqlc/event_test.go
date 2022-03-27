package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/heritechie/motiket/api/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	arg := CreateEventParams{
		ID:   uuid.New(),
		Name: util.RandomString(12),
		Description: sql.NullString{
			String: util.RandomString(255),
			Valid:  true,
		},
		Prefix: util.RandomString(4),
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.ID, event.ID)
	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.Description, event.Description)
	require.Equal(t, arg.Prefix, event.Prefix)

	require.NotZero(t, event.CreatedAt)
	require.NotZero(t, event.UpdatedAt)

	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestGetEvent(t *testing.T) {
	account1 := createRandomEvent(t)
	account2, err := testQueries.GetEvent(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Description, account2.Description)
	require.Equal(t, account1.Prefix, account2.Prefix)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	require.WithinDuration(t, account1.UpdatedAt, account2.UpdatedAt, time.Second)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	arg := ListEventParams{
		Limit:  5,
		Offset: 5,
	}

	events, err := testQueries.ListEvent(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, events, 5)

	for _, event := range events {
		require.NotEmpty(t, event)
	}
}
