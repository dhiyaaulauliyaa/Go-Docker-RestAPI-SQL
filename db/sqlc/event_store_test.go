package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEventTx(t *testing.T) {
	store := NewStore(testDB)

	n := 5
	errs := make(chan error)
	results := make(chan CreateEventTxResult)

	for i := 0; i < n; i++ {
		date, _ := time.Parse("2006-01-02", "2020-01-29")

		var params = CreateEventTxParams{
			Name:   "Test Cr Event Tx",
			Venue:  "Masjid",
			Masjid: sql.NullInt32{},
			Date:   date,
		}

		go func() {
			res, err := store.CreateEventTx(context.Background(), params)
			errs <- err
			results <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		res := <-results
		var event = res.Event

		date, _ := time.Parse("2006-01-02", "2020-01-29")

		var params = CreateEventTxParams{
			Name:   "Test Cr Event Tx",
			Venue:  "Masjid",
			Masjid: sql.NullInt32{},
			Date:   date,
		}

		require.NoError(t, err)
		require.NotEmpty(t, event)

		require.Equal(t, params.Name, event.Name)
		require.Equal(t, params.Venue, event.Venue)
		// require.Equal(t, sql.NullInt32{}, event.Masjid)

		require.NotZero(t, event.ID)
		require.NotZero(t, event.CreatedAt)

	}

}

func TestDeleteEventTx(t *testing.T) {
	store := NewStore(testDB)

	n := 5
	errs := make(chan error)
	results := make(chan DeleteEventTxResult)

	for i := 0; i < n; i++ {
		go func() {
			testEvent := createRandomEvent(t)

			/* Get event */
			res, err := store.DeleteEventTx(context.Background(), testEvent.ID)
			errs <- err
			results <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		res := <-results

		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Contains(t, res.Message, "Success")
	}
}
