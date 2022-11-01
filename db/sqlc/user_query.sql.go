// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user_query.sql

package db

import (
	"context"
	"database/sql"
)

const changePassword = `-- name: ChangePassword :one
UPDATE
  users
SET
  password = $2
WHERE
  id = $1 RETURNING id, username, password, name, phone, email, gender, age, avatar, created_at
`

type ChangePasswordParams struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, changePassword, arg.ID, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Age,
		&i.Avatar,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
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
  ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, username, password, name, phone, email, gender, age, avatar, created_at
`

type CreateUserParams struct {
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Phone    string         `json:"phone"`
	Email    sql.NullString `json:"email"`
	Gender   int32          `json:"gender"`
	Age      int32          `json:"age"`
	Avatar   sql.NullString `json:"avatar"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Name,
		arg.Password,
		arg.Phone,
		arg.Email,
		arg.Gender,
		arg.Age,
		arg.Avatar,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Age,
		&i.Avatar,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM
  users
WHERE
  id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
  id, username, password, name, phone, email, gender, age, avatar, created_at
FROM
  users
WHERE
  id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Age,
		&i.Avatar,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT
  id, username, password, name, phone, email, gender, age, avatar, created_at
FROM
  users
ORDER BY
  name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.Phone,
			&i.Email,
			&i.Gender,
			&i.Age,
			&i.Avatar,
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

const updateUser = `-- name: UpdateUser :one
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
  id = $1 RETURNING id, username, password, name, phone, email, gender, age, avatar, created_at
`

type UpdateUserParams struct {
	ID       int32          `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Phone    string         `json:"phone"`
	Email    sql.NullString `json:"email"`
	Gender   int32          `json:"gender"`
	Age      int32          `json:"age"`
	Avatar   sql.NullString `json:"avatar"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Name,
		arg.Phone,
		arg.Email,
		arg.Gender,
		arg.Age,
		arg.Avatar,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Phone,
		&i.Email,
		&i.Gender,
		&i.Age,
		&i.Avatar,
		&i.CreatedAt,
	)
	return i, err
}
