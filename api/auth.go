package api

import (
	"errors"

	"github.com/dhiyaaulauliyaa/learn-go/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) checkAuthOwnership(ctx *gin.Context, identifier string) error {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if identifier != authPayload.Identifier {
		return errors.New("account doesn't belong to the authenticated user")
	}

	return nil
}
