-- name: GetEvent :one
SELECT
  *
FROM
  events
WHERE
  id = $1;


-- name: ListEvents :many
SELECT
  *
FROM
  events
ORDER BY
  name;


-- name: CreateEvent :one
INSERT INTO
  events (name, venue, masjid, date)
VALUES
  ($1, $2, $3, $4) RETURNING *;


-- name: DeleteEvent :exec
DELETE FROM
  events
WHERE
  id = $1;


-- name: UpdateEvent :one
UPDATE
  events
SET
  name = $2,
  venue = $3,
  masjid = $4,
  date = $5
WHERE
  id = $1 RETURNING *;