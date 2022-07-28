-- name: GetUser :one
SELECT
  *
FROM
  users
WHERE
  id = $1;


-- name: ListUsers :many
SELECT
  *
FROM
  users
ORDER BY
  name;


-- name: CreateUser :one
INSERT INTO
  users (username, name, phone, gender, age, avatar)
VALUES
  ($1, $2, $3, $4, $5, $6) RETURNING *;


-- name: DeleteUser :exec
DELETE FROM
  users
WHERE
  id = $1;


-- name: UpdateUser :one
UPDATE
  users
SET
  username = $2,
  name = $3,
  phone = $4,
  gender = $5,
  age = $6,
  avatar = $7
WHERE
  id = $1 RETURNING *;