package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	util "github.com/dhiyaaulauliyaa/learn-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	date, _ := time.Parse("2006-01-02", "2020-01-29")

	arg := CreateEventParams{
		Name:   fmt.Sprintf("Event %s", util.RandomString(5)),
		Venue:  fmt.Sprintf("Masjid Al-%s", util.RandomString(5)),
		Masjid: sql.NullInt32{},
		Date:   date,
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.Venue, event.Venue)
	require.Equal(t, arg.Masjid, event.Masjid)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)

	return event
}

func TestCreateEvent(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2020-01-29")

	arg := CreateEventParams{
		Name:   "Test Event",
		Venue:  "Masjid",
		Masjid: sql.NullInt32{},
		Date:   date,
	}

	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.Venue, event.Venue)
	require.Equal(t, arg.Masjid, event.Masjid)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)
}

func TestGetEvent(t *testing.T) {
	testEvent := createRandomEvent(t)

	event, err := testQueries.GetEvent(context.Background(), testEvent.ID)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, testEvent.Name, event.Name)
	require.Equal(t, testEvent.Venue, event.Venue)
	require.Equal(t, testEvent.Masjid, event.Masjid)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)
	require.WithinDuration(t, testEvent.CreatedAt, event.CreatedAt, time.Second)
}

func TestListEvents(t *testing.T) {
	createRandomEvent(t)
	createRandomEvent(t)

	event, err := testQueries.ListEvents(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, event)
}

func TestUpdateEvent(t *testing.T) {
	testEvent := createRandomEvent(t)

	date, _ := time.Parse("2006-01-02", "2020-01-29")

	arg := UpdateEventParams{
		ID:     testEvent.ID,
		Name:   "Test Update Event",
		Venue:  "Masjid",
		Masjid: sql.NullInt32{},
		Date:   date,
	}

	event, err := testQueries.UpdateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.Equal(t, arg.Name, event.Name)
	require.Equal(t, arg.Venue, event.Venue)
	require.Equal(t, arg.Masjid, event.Masjid)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.CreatedAt)

}

func TestDeleteEvent(t *testing.T) {
	testEvent := createRandomEvent(t)

	err := testQueries.DeleteEvent(context.Background(), testEvent.ID)
	require.NoError(t, err)

}
