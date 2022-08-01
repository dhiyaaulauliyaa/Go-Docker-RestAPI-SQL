package api

import (
	"database/sql"
	"net/http"

	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	util "github.com/dhiyaaulauliyaa/learn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

func userErrHandling(err error, defaultMsg string) (string, int) {
	if err == sql.ErrNoRows {
		return "Data not found", http.StatusNotFound
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			switch pqErr.Constraint {
			case "users_email_key":
				return "Email already used", http.StatusForbidden
			case "users_username_key":
				return "Username already used", http.StatusForbidden
			case "users_phone_key":
				return "Phone number already used", http.StatusForbidden
			}
		}
	}

	return defaultMsg, http.StatusInternalServerError
}

type createUserRequest struct {
	Username string      `json:"username" binding:"required,alphanum"`
	Name     string      `json:"name" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Phone    string      `json:"phone" binding:"required"`
	Email    null.String `json:"email" binding:"required,email"`
	Gender   int32       `json:"gender" binding:"required"`
	Age      int32       `json:"age" binding:"required"`
	Avatar   null.String `json:"avatar"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
		return
	}

	/* Hash password */
	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorHashPass))
		return
	}

	arg := db.CreateUserTxParams{
		Username: req.Username,
		Name:     req.Name,
		Password: hashedPass,
		Phone:    req.Phone,
		Email:    util.NullableToString(req.Email),
		Gender:   req.Gender,
		Age:      req.Age,
		Avatar:   util.NullableToString(req.Avatar),
	}

	res, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		message, code := userErrHandling(err, "Create user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindUri))
		return
	}

	res, err := server.store.GetUserTx(ctx, req.ID)
	if err != nil {
		message, code := userErrHandling(err, "Get user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) getUsers(ctx *gin.Context) {
	res, err := server.store.GetUsersTx(ctx)
	if err != nil {
		message, code := userErrHandling(err, "Get users failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type updateUserRequest struct {
	ID       int32       `json:"id" binding:"required"`
	Username string      `json:"username" binding:"required"`
	Name     string      `json:"name" binding:"required"`
	Phone    string      `json:"phone" binding:"required"`
	Email    null.String `json:"email" binding:"email"`
	Gender   int32       `json:"gender" binding:"required"`
	Age      int32       `json:"age" binding:"required"`
	Avatar   null.String `json:"avatar"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
		return
	}

	arg := db.UpdateUserTxParams{
		ID:       req.ID,
		Username: req.Username,
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    util.NullableToString(req.Email),
		Gender:   req.Gender,
		Age:      req.Age,
		Avatar:   util.NullableToString(req.Avatar),
	}

	res, err := server.store.UpdateUserTx(ctx, arg)
	if err != nil {
		message, code := userErrHandling(err, "Update user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type deleteUserRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindUri))
		return
	}

	res, err := server.store.DeleteUserTx(ctx, req.ID)
	if err != nil {
		message, code := userErrHandling(err, "Delete user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
