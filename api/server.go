package api

import (
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	errorBindUri  = "Param parsing failed"
	errorBindBody = "Body parsing failed"
	errorHashPass = "Password hashing failed"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/event", server.createEvent)
	router.GET("/events", server.getEvents)
	router.GET("/event/:id", server.getEvent)
	router.PUT("/event", server.updateEvent)
	router.DELETE("/event/:id", server.deleteEvent)

	router.POST("/user", server.createUser)
	router.GET("/users", server.getUsers)
	router.GET("/user/:id", server.getUser)
	router.PUT("/user", server.updateUser)
	router.DELETE("/user/:id", server.deleteUser)

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
