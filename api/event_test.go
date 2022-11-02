package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/dhiyaaulauliyaa/learn-go/db/mock"
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/dhiyaaulauliyaa/learn-go/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func randomEvent() db.Event {
	return db.Event{
		ID:        int32(util.RandomAge()),
		Name:      util.RandomName(),
		Venue:     util.RandomString(5),
		Community: sql.NullInt32{},
		Masjid:    sql.NullInt32{},
		Date:      util.RandomDate(),
		CreatedAt: util.RandomDate(),
	}
}

func TestCreateEventAPI(t *testing.T) {
	/* Arrange Param */
	event := randomEvent()
	arg := db.CreateEventParams{
		Name:   event.Name,
		Venue:  event.Venue,
		Masjid: event.Masjid,
		Date:   event.Date,
	}

	/* Test Cases */
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":      arg.Name,
				"venue":     arg.Venue,
				"community": null.Int{},
				"masjid":    null.Int{},
				"date":      arg.Date,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateEvent(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(event, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				checkGetEventResponseBody(t, recorder.Body, event)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateEvent(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"name":      arg.Name,
				"venue":     arg.Venue,
				"community": null.Int{},
				"masjid":    null.Int{},
				"date":      arg.Date,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateEvent(gomock.Any(), gomock.Eq(arg)).
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

		t.Run(
			tc.name,
			func(t *testing.T) {
				/* Arrange */
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)
				tc.buildStubs(store)

				server, err := NewServer(store)
				require.NoError(t, err)

				recorder := httptest.NewRecorder()

				/* Covert body to json */
				data, err := json.Marshal(tc.body)
				require.NoError(t, err)

				/* Act */
				url := "/event"
				req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, req)

				/* Assert */
				tc.checkResponse(t, recorder)
			},
		)
	}
}

func TestGetEventAPI(t *testing.T) {
	/* Arrange Param */
	event := randomEvent()

	/* Test Cases */
	testCases := []struct {
		name          string
		eventID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
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
				checkGetEventResponseBody(t, recorder.Body, event)
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
			name:    "InternalServerError",
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
		{
			name:    "BadRequest",
			eventID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEvent(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				/* Arrange */
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)
				tc.buildStubs(store)

				server, err := NewServer(store)
				require.NoError(t, err)

				recorder := httptest.NewRecorder()

				/* Act */
				url := fmt.Sprintf("/event/%d", tc.eventID)
				req, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, req)

				/* Assert */
				tc.checkResponse(t, recorder)
			},
		)
	}
}

func TestGetEventsAPI(t *testing.T) {
	/* Arrange Param */
	n := 5
	events := make([]db.Event, n)
	for i := 0; i < n; i++ {
		events[i] = randomEvent()
	}

	/* Test Cases */
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListEvents(gomock.Any()).
					Times(1).
					Return(events, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				checkGetEventsResponseBody(t, recorder.Body, events)
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListEvents(gomock.Any()).
					Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				/* Arrange */
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)
				tc.buildStubs(store)

				server, err := NewServer(store)
				require.NoError(t, err)

				recorder := httptest.NewRecorder()

				/* Act */
				url := "/events"
				req, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)
				server.router.ServeHTTP(recorder, req)

				/* Assert */
				tc.checkResponse(t, recorder)
			},
		)
	}
}

func checkGetEventResponseBody(t *testing.T, body *bytes.Buffer, event db.Event) {
	/* Read body from buffer */
	rawResp, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	/* Convert body from json to Response */
	var response Response
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		panic(err)
	}

	/* Get Event from Response */
	eventData := getEventFromResponse(response)

	/* Check */
	require.NoError(t, err)
	require.Equal(t, event, eventData)
}
func getEventFromResponse(response Response) db.Event {
	/* Convert data from any to json */
	data, err := json.Marshal(response.Data)
	if err != nil {
		panic(err)
	}

	/* Convert data from json to Event */
	var event db.Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		panic(err)
	}

	return event
}

func checkGetEventsResponseBody(t *testing.T, body *bytes.Buffer, events []db.Event) {
	/* Read body from buffer */
	rawResp, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	/* Convert body from json to Response */
	var response Response
	err = json.Unmarshal(rawResp, &response)
	if err != nil {
		panic(err)
	}

	/* Get Event from Response */
	eventData := getEventsFromResponse(response)

	/* Check */
	require.NoError(t, err)
	require.Equal(t, events, eventData)
}

func getEventsFromResponse(response Response) []db.Event {
	/* Convert data from any to json */
	data, err := json.Marshal(response.Data)
	if err != nil {
		panic(err)
	}

	/* Convert data from json to Event */
	var events []db.Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		panic(err)
	}

	return events
}
