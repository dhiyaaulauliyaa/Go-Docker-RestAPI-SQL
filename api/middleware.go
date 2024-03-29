package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
// func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

// 		if len(authorizationHeader) == 0 {
// 			msg := "authorization header is not provided"
// 			err := errors.New(msg)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
// 			return
// 		}

// 		fields := strings.Fields(authorizationHeader)
// 		if len(fields) < 2 {
// 			msg := "invalid authorization header format"
// 			err := errors.New(msg)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
// 			return
// 		}

// 		authorizationType := strings.ToLower(fields[0])
// 		if authorizationType != authorizationTypeBearer {
// 			msg := "unsupported authorization type"
// 			err := fmt.Errorf("%s: %s", msg, authorizationType)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
// 			return
// 		}

// 		accessToken := fields[1]
// 		payload, err := tokenMaker.VerifyToken(accessToken)
// 		if err != nil {
// 			msg := "token is invalid"
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
// 			return
// 		}

// 		ctx.Set(authorizationPayloadKey, payload)
// 		ctx.Next()
// 	}
// }

func firebaseAuthMiddleware(firebaseAuth auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("firebaseAuth", firebaseAuth)

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			msg := "authorization header is not provided"
			err := errors.New(msg)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			msg := "invalid authorization header format"
			err := errors.New(msg)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			msg := "unsupported authorization type"
			err := fmt.Errorf("%s: %s", msg, authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
			return
		}

		/* Verify token */
		accessToken := fields[1]
		token, err := firebaseAuth.VerifyIDToken(context.Background(), accessToken)
		if err != nil {
			msg := "token is invalid"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err, msg))
			return
		}

		ctx.Set("UUID", token.UID)
		ctx.Next()
	}
}
