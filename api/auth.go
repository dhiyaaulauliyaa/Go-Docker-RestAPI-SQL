package api

import (
	"errors"
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func (server *Server) checkAuthOwnership(ctx *gin.Context, identifier string) error {
	auth := ctx.MustGet("firebaseAuth").(auth.Client)
	_, err := auth.GetUserByPhoneNumber(ctx, fmt.Sprint("+", identifier))

	if err != nil {
		fmt.Println(err)
		return errors.New("account doesn't belong to the authenticated user")
	}

	return nil
}

// type refreshTokenReq struct {
// 	RefreshToken string `json:"refresh_token" binding:"required"`
// }

// type refreshTokenRes struct {
// 	AccessToken        string    `json:"access_token"`
// 	AccessTokenExpDate time.Time `json:"access_token_expires_at"`
// }

// func (server *Server) refreshToken(ctx *gin.Context) {
// 	var req refreshTokenReq

// 	/* Parse body to get refresh token */
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
// 		return
// 	}

// 	/* Verify the refresh token */
// 	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, "token invalid"))
// 		return
// 	}

// 	/* Get token session */
// 	session, err := server.store.GetSession(ctx, refreshPayload.ID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err, "token not found"))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "session invalid"))
// 		return
// 	}

// 	/* Check token validity based on session */
// 	if session.IsBlocked {
// 		msg := "session has been blocked"
// 		err := fmt.Errorf(msg)
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, msg))
// 		return
// 	}

// 	log.Println(session.Phone)
// 	log.Println(refreshPayload.Identifier)

// 	if session.Phone != refreshPayload.Identifier {
// 		msg := "incorrect user session"
// 		err := fmt.Errorf(msg)
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, msg))
// 		return
// 	}

// 	if session.RefreshToken != req.RefreshToken {
// 		msg := "mismatched session token"
// 		err := fmt.Errorf(msg)
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, msg))
// 		return
// 	}

// 	if time.Now().After(session.ExpiresAt) {
// 		msg := "session has been expired"
// 		err := fmt.Errorf(msg)
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, msg))
// 		return
// 	}

// 	/* Create new access token */
// 	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
// 		refreshPayload.Identifier,
// 		time.Duration(60*15*1000000000),
// 	)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}
// 	if accessPayload == nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, "Create token failed"))
// 		return
// 	}

// 	res := refreshTokenRes{
// 		AccessToken:        accessToken,
// 		AccessTokenExpDate: accessPayload.ExpiredAt,
// 	}
// 	ctx.JSON(http.StatusOK, res)
// }
