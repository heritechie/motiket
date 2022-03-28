package http

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/heritechie/motiket/api/internal/db/mock"
	db "github.com/heritechie/motiket/api/internal/db/sqlc"
	"github.com/heritechie/motiket/api/internal/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountHTTP(t *testing.T) {
	event := randomEvent()

	testCases := []struct {
		name          string
		eventID       uuid.UUID
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			eventID: event.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEvent(gomock.Any(), gomock.Eq(event.ID)).
					Times(1).
					Return(event, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEvent(t, recorder.Body, event)
			},
		},
		{
			name:    "NotFound",
			eventID: event.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEvent(gomock.Any(), gomock.Eq(event.ID)).
					Times(1).
					Return(db.Event{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			eventID: event.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEvent(gomock.Any(), gomock.Eq(event.ID)).
					Times(1).
					Return(db.Event{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/events/%s", tc.eventID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}

}

func randomEvent() db.Event {
	return db.Event{
		ID:   uuid.New(),
		Name: util.RandomString(8),
		Description: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
		Prefix: util.RandomString(3),
	}
}

func requireBodyMatchEvent(t *testing.T, body *bytes.Buffer, event db.Event) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getEvent db.Event
	err = json.Unmarshal(data, &getEvent)
	require.NoError(t, err)
	require.Equal(t, event, getEvent)
}
