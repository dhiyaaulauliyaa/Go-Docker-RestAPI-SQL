package api

import (
	"firebase.google.com/go/auth"
	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	errorBindUri  = "Param parsing failed"
	errorBindBody = "Body parsing failed"
	errorHashPass = "Password hashing failed"
)

type Server struct {
	store        db.Store
	firebaseAuth auth.Client

	router *gin.Engine
}

func NewServer(store db.Store, firebaseAuth auth.Client) (*Server, error) {
	server := &Server{
		store:        store,
		firebaseAuth: firebaseAuth,
	}

	server.initRouter()

	return server, nil
}

func (server *Server) initRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(firebaseAuthMiddleware(server.firebaseAuth))

	// router.POST("/auth/refresh_token", server.refreshToken)

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
