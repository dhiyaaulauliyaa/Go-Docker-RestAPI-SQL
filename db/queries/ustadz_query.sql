-- name: GetUstadz :one
SELECT
  *
FROM
  ustadzs
WHERE
  id = $1;


-- name: ListUstadz :many
SELECT
  *
FROM
  ustadzs
ORDER BY
  name;


-- name: CreateUstadz :one
INSERT INTO
  ustadzs (name, avatar, gender)
VALUES
  ($1, $2, $3) RETURNING *;


-- name: DeleteUstadz :exec
DELETE FROM
  ustadzs
WHERE
  id = $1;


-- name: UpdateUstadz :one
UPDATE
  ustadzs
SET
  name = $2,
  avatar = $3,
  gender = $4
WHERE
  id = $1 RETURNING *;