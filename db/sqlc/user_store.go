package db

import (
	"context"
	"database/sql"

	util "github.com/dhiyaaulauliyaa/learn-go/utils"
	"gopkg.in/guregu/null.v4"
)

type UserResponse struct {
	ID       int32       `json:"id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Phone    string      `json:"phone"`
	Email    null.String `json:"email"`
	Gender   int32       `json:"gender"`
	Age      int32       `json:"age"`
	Avatar   null.String `json:"avatar"`
}

func generateUserResponse(user User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    util.StringToNullable(user.Email),
		Gender:   user.Gender,
		Age:      user.Age,
		Avatar:   util.StringToNullable(user.Avatar),
	}
}

type CreateUserTxParams struct {
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Phone    string         `json:"phone"`
	Email    sql.NullString `json:"email"`
	Gender   int32          `json:"gender"`
	Age      int32          `json:"age"`
	Avatar   sql.NullString `json:"avatar"`
}

type CreateUserTxResult struct {
	Message string       `json:"message"`
	User    UserResponse `json:"data"`
}

func (store *Store) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	/* Create User */
	err := store.execTx(ctx, func(q *Queries) error {
		user, err := q.CreateUser(ctx, CreateUserParams(arg))
		if err != nil {
			return err
		}

		result.User = generateUserResponse(user)
		return nil
	})

	if err == nil {
		result.Message = "Success create user"
	}
	return result, err
}

type GetUserTxResult struct {
	Message string       `json:"message"`
	User    UserResponse `json:"data"`
}

func (store *Store) GetUserTx(ctx context.Context, id int32) (GetUserTxResult, error) {
	var result GetUserTxResult

	/* Get User */
	err := store.execTx(ctx, func(q *Queries) error {
		user, err := q.GetUser(ctx, id)
		if err != nil {
			return err
		}

		result.User = generateUserResponse(user)
		return nil
	})

	if err == nil {
		result.Message = "Success get user"
	}
	return result, err
}

type ListUsersTxResult struct {
	Message string         `json:"message"`
	User    []UserResponse `json:"data"`
}

func (store *Store) GetUsersTx(ctx context.Context) (ListUsersTxResult, error) {
	var result ListUsersTxResult

	/* Get User */
	err := store.execTx(ctx, func(q *Queries) error {
		users, err := q.ListUsers(ctx)
		if err != nil {
			return err
		}

		var usersRes []UserResponse
		for _, user := range users {
			usersRes = append(usersRes, generateUserResponse(user))
		}

		result.User = usersRes
		return nil
	})

	if err == nil {
		result.Message = "Success get users"
	}
	return result, err
}

type UpdateUserTxParams struct {
	ID       int32          `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Phone    string         `json:"phone"`
	Email    sql.NullString `json:"email"`
	Gender   int32          `json:"gender"`
	Age      int32          `json:"age"`
	Avatar   sql.NullString `json:"avatar"`
}

type UpdateUserTxResult struct {
	Message string       `json:"message"`
	User    UserResponse `json:"data"`
}

func (store *Store) UpdateUserTx(ctx context.Context, arg UpdateUserTxParams) (UpdateUserTxResult, error) {
	var result UpdateUserTxResult

	/* Update User */
	err := store.execTx(ctx, func(q *Queries) error {
		user, err := q.UpdateUser(ctx, UpdateUserParams(arg))
		if err != nil {
			return err
		}

		result.User = generateUserResponse(user)
		return nil
	})

	if err == nil {
		result.Message = "Success update user"
	}
	return result, err
}

type DeleteUserTxResult struct {
	Message string `json:"message"`
}

func (store *Store) DeleteUserTx(ctx context.Context, id int32) (DeleteUserTxResult, error) {
	var result DeleteUserTxResult
	var err error

	/* Check if data exist */
	err = store.execTx(ctx, func(q *Queries) error {
		_, err := q.GetUser(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return result, err
	}

	/* Start delete data */
	err = store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteUser(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		result.Message = "Success delete user"
	}
	return result, err
}
