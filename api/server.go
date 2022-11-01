package api

import (
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	errorBindUri  = "Fail to parse id"
	errorBindBody = "Fail to parse body"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/event/:id", server.getEvent)
	router.GET("/events", server.getEvents)
	router.POST("/event", server.createEvent)
	router.PUT("/event", server.updateEvent)
	router.DELETE("/event/:id", server.deleteEvent)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error, message string) gin.H {
	return gin.H{
		"message": message,
		"error":   err.Error(),
	}
}
