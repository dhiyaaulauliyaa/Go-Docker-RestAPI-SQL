-- name: GetMasjid :one
SELECT
  *
FROM
  masjids
WHERE
  id = $1;


-- name: ListMasjids :many
SELECT
  *
FROM
  masjids
ORDER BY
  name;


-- name: CreateMasjid :one
INSERT INTO
  masjids (name, address, city, coordinate, logo)
VALUES
  ($1, $2, $3, $4, $5) RETURNING *;


-- name: DeleteMasjid :exec
DELETE FROM
  masjids
WHERE
  id = $1;


-- name: UpdateMasjid :one
UPDATE
  masjids
SET
  name = $2,
  address = $3,
  city = $4,
  coordinate = $5,
  logo = $6
WHERE
  id = $1 RETURNING *;