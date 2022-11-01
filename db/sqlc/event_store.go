package db

import (
	"context"
	"database/sql"
	"time"

	util "github.com/dhiyaaulauliyaa/learn-go/util"
	"gopkg.in/guregu/null.v4"
)

type EventResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Venue     string    `json:"venue"`
	Masjid    null.Int  `json:"masjid"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt"`
}

func generateEventResponse(event Event) EventResponse {
	return EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Venue:     event.Venue,
		Masjid:    util.Int32ToNullable(event.Masjid),
		Date:      event.Date,
		CreatedAt: event.CreatedAt,
	}
}

type CreateEventTxParams struct {
	Name   string        `json:"name"`
	Venue  string        `json:"venue"`
	Masjid sql.NullInt32 `json:"masjid"`
	Date   time.Time     `json:"date"`
}

type CreateEventTxResult struct {
	Message string        `json:"message"`
	Event   EventResponse `json:"data"`
}

func (store *Store) CreateEventTx(ctx context.Context, arg CreateEventTxParams) (CreateEventTxResult, error) {
	var result CreateEventTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		event, err := q.CreateEvent(ctx, CreateEventParams(arg))
		if err != nil {
			return err
		}

		result.Event = generateEventResponse(event)
		return nil
	})

	if err == nil {
		result.Message = "Success create event"
	}
	return result, err
}

type GetEventTxResult struct {
	Message string        `json:"message"`
	Event   EventResponse `json:"data"`
}

func (store *Store) GetEventTx(ctx context.Context, id int32) (GetEventTxResult, error) {
	var result GetEventTxResult

	/* Get Event */
	err := store.execTx(ctx, func(q *Queries) error {
		event, err := q.GetEvent(ctx, id)
		if err != nil {
			return err
		}

		result.Event = generateEventResponse(event)
		return nil
	})

	if err == nil {
		result.Message = "Success get event"
	}
	return result, err
}

type ListEventsTxResult struct {
	Message string          `json:"message"`
	Event   []EventResponse `json:"data"`
}

func (store *Store) GetEventsTx(ctx context.Context) (ListEventsTxResult, error) {
	var result ListEventsTxResult

	/* Get Event */
	err := store.execTx(ctx, func(q *Queries) error {
		events, err := q.ListEvents(ctx)
		if err != nil {
			return err
		}

		var eventsRes []EventResponse
		for _, event := range events {
			eventsRes = append(eventsRes, generateEventResponse(event))
		}

		result.Event = eventsRes
		return nil
	})

	if err == nil {
		result.Message = "Success get events"
	}
	return result, err
}

type UpdateEventTxParams struct {
	ID     int32         `json:"id"`
	Name   string        `json:"name"`
	Venue  string        `json:"venue"`
	Masjid sql.NullInt32 `json:"masjid"`
	Date   time.Time     `json:"date"`
}

type UpdateEventTxResult struct {
	Message string        `json:"message"`
	Event   EventResponse `json:"data"`
}

func (store *Store) UpdateEventTx(ctx context.Context, arg UpdateEventTxParams) (UpdateEventTxResult, error) {
	var result UpdateEventTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		event, err := q.UpdateEvent(ctx, UpdateEventParams(arg))
		if err != nil {
			return err
		}

		result.Event = generateEventResponse(event)
		return nil
	})

	if err == nil {
		result.Message = "Success get events"
	}
	return result, err
}

type DeleteEventTxResult struct {
	Message string `json:"message"`
}

func (store *Store) DeleteEventTx(ctx context.Context, id int32) (DeleteEventTxResult, error) {
	var result DeleteEventTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteEvent(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		result.Message = "Success delete events"
	}
	return result, err
}
