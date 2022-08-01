package api

import (
	"fmt"

	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/dhiyaaulauliyaa/learn-go/token"
	"github.com/gin-gonic/gin"
)

const (
	errorBindUri  = "Param parsing failed"
	errorBindBody = "Body parsing failed"
	errorHashPass = "Password hashing failed"

	tokenSymmetricKey   = "12345678901234567890123456789012"
	accessTokenDuration = "15m"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("token maker creation failed: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.initRouter()

	return server, nil
}

func (server *Server) initRouter() {
	router := gin.Default()

	router.POST("/event", server.createEvent)
	router.GET("/events", server.getEvents)
	router.GET("/event/:id", server.getEvent)
	router.PUT("/event", server.updateEvent)
	router.DELETE("/event/:id", server.deleteEvent)

	router.POST("/user/login", server.userLogin)
	router.POST("/user", server.createUser)
	router.GET("/users", server.getUsers)
	router.GET("/user/:id", server.getUser)
	router.PUT("/user", server.updateUser)
	router.DELETE("/user/:id", server.deleteUser)

	server.router = router
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
