// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: event_query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO
  events (name, venue, masjid, date)
VALUES
  ($1, $2, $3, $4) RETURNING id, name, venue, masjid, date, created_at
`

type CreateEventParams struct {
	Name   string        `json:"name"`
	Venue  string        `json:"venue"`
	Masjid sql.NullInt32 `json:"masjid"`
	Date   time.Time     `json:"date"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Name,
		arg.Venue,
		arg.Masjid,
		arg.Date,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Venue,
		&i.Masjid,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM
  events
WHERE
  id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEvent = `-- name: GetEvent :one
SELECT
  id, name, venue, masjid, date, created_at
FROM
  events
WHERE
  id = $1
`

func (q *Queries) GetEvent(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Venue,
		&i.Masjid,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT
  id, name, venue, masjid, date, created_at
FROM
  events
ORDER BY
  name
`

func (q *Queries) ListEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Venue,
			&i.Masjid,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
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

const updateEvent = `-- name: UpdateEvent :one
UPDATE
  events
SET
  name = $2,
  venue = $3,
  masjid = $4,
  date = $5
WHERE
  id = $1 RETURNING id, name, venue, masjid, date, created_at
`

type UpdateEventParams struct {
	ID     int32         `json:"id"`
	Name   string        `json:"name"`
	Venue  string        `json:"venue"`
	Masjid sql.NullInt32 `json:"masjid"`
	Date   time.Time     `json:"date"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEvent,
		arg.ID,
		arg.Name,
		arg.Venue,
		arg.Masjid,
		arg.Date,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Venue,
		&i.Masjid,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}