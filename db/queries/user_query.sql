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
  users (
    username,
    name,
    password,
    phone,
    email,
    gender,
    age,
    avatar
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;


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
  email = $5,
  gender = $6,
  age = $7,
  avatar = $8
WHERE
  id = $1 RETURNING *;


-- name: ChangePassword :one
UPDATE
  users
SET
  password = $2
WHERE
  id = $1 RETURNING *;