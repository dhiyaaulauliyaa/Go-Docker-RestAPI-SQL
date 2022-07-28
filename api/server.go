package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Server {
	server := &Server{db: db}
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

func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}
