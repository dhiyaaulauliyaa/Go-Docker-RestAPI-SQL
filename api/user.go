package api

import (
	"database/sql"
	"net/http"

	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	util "github.com/dhiyaaulauliyaa/learn-go/util"
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

	arg := db.CreateUserParams{
		Username: req.Username,
		Name:     req.Name,
		Password: hashedPass,
		Phone:    req.Phone,
		Email:    util.NullableToString(req.Email),
		Gender:   req.Gender,
		Age:      req.Age,
		Avatar:   util.NullableToString(req.Avatar),
	}

	res, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		message, code := userErrHandling(err, "Create user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
}

type getUserRequest struct {
	Phone string `uri:"phone" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	/* Parse URI to get ID */
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindUri))
		return
	}

	/* Start get User */
	res, err := server.store.GetUser(ctx, req.Phone)
	if err != nil {
		message, code := userErrHandling(err, "Get user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	/* Check auth */
	err = server.checkAuthOwnership(ctx, res.Phone)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
}

func (server *Server) getUsers(ctx *gin.Context) {
	res, err := server.store.ListUsers(ctx)
	if err != nil {
		message, code := userErrHandling(err, "Get users failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
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

	/* Check auth */
	err := server.checkAuthOwnership(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, err.Error()))
		return
	}

	/* Process args */
	arg := db.UpdateUserParams{
		ID:       req.ID,
		Username: req.Username,
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    util.NullableToString(req.Email),
		Gender:   req.Gender,
		Age:      req.Age,
		Avatar:   util.NullableToString(req.Avatar),
	}

	/* Start Update */
	res, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		message, code := userErrHandling(err, "Update user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
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

	err := server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		message, code := userErrHandling(err, "Delete user failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(nil))
}

// type userLoginRequest struct {
// 	Phone    string `json:"phone" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// type userLoginResponse struct {
// 	SessionID           uuid.UUID `json:"session_id"`
// 	Message             string    `json:"message"`
// 	AccessToken         string    `json:"access_token"`
// 	RefreshToken        string    `json:"refresh_token"`
// 	AccessTokenExpDate  time.Time `json:"access_token_exp_date"`
// 	RefreshTokenExpDate time.Time `json:"refresh_token_exp_date"`
// 	User                db.User   `json:"user"`
// }

// func (server *Server) userLogin(ctx *gin.Context) {
// 	var req userLoginRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
// 		return
// 	}

// 	/* Get User */
// 	user, err := server.store.GetUser(ctx, req.Phone)
// 	if err != nil {
// 		message, code := userErrHandling(err, "Get user failed")
// 		ctx.JSON(code, errorResponse(err, message))
// 		return
// 	}

// 	/* Check Password */
// 	err = util.CheckPassword(req.Password, user.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, "Password incorrect"))
// 		return
// 	}

// 	/* Generate Access Token */
// 	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
// 		user.Phone,
// 		time.Duration(15*60*1000000000),
// 	)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}
// 	if accessPayload == nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}

// 	/* Generate Refresh Token */
// 	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
// 		user.Phone,
// 		time.Duration(30*24*60*60*1000000000),
// 	)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}
// 	if refreshPayload == nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}

// 	log.Println(refreshPayload.Identifier)

// 	/* Store User Session */
// 	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
// 		ID:           refreshPayload.ID,
// 		Phone:        user.Phone,
// 		RefreshToken: refreshToken,
// 		UserAgent:    ctx.Request.UserAgent(),
// 		ClientIp:     ctx.ClientIP(),
// 		IsBlocked:    false,
// 		ExpiresAt:    refreshPayload.ExpiredAt,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create session error"))
// 		return
// 	}

// 	var res = userLoginResponse{
// 		SessionID:           session.ID,
// 		AccessToken:         accessToken,
// 		RefreshToken:        refreshToken,
// 		AccessTokenExpDate:  accessPayload.ExpiredAt,
// 		RefreshTokenExpDate: refreshPayload.ExpiredAt,
// 		User:                user,
// 	}

// 	ctx.JSON(http.StatusOK, successResponse(res))
// }
