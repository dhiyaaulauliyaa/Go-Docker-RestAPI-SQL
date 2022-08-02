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

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/auth/refresh_token", server.refreshToken)

	router.POST("/user/login", server.userLogin)
	router.POST("/user", server.createUser)
	authRoutes.GET("/users", server.getUsers)
	authRoutes.GET("/user/:phone", server.getUser)
	authRoutes.PUT("/user", server.updateUser)
	authRoutes.DELETE("/user/:id", server.deleteUser)

	router.POST("/event", server.createEvent)
	router.GET("/events", server.getEvents)
	router.GET("/event/:id", server.getEvent)
	authRoutes.PUT("/event", server.updateEvent)
	authRoutes.DELETE("/event/:id", server.deleteEvent)

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
