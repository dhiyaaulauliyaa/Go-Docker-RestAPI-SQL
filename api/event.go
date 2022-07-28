package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEventRequest struct {
	Name   string        `json:"name" binding:"required"`
	Venue  string        `json:"venue" binding:"required"`
	Masjid sql.NullInt32 `json:"masjid"`
	Date   time.Time     `json:"date" binding:"required"`
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEventParams{
		Name:   req.Name,
		Venue:  req.Venue,
		Masjid: req.Masjid,
		Date:   req.Date,
	}

	q := db.New(server.db)
	events, err := q.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, events)
}

type getEventRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q := db.New(server.db)
	events, err := q.GetEvent(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (server *Server) getEvents(ctx *gin.Context) {
	q := db.New(server.db)
	events, err := q.ListEvents(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, events)
}

type updateEventRequest struct {
	ID     int32         `json:"id" binding:"required"`
	Name   string        `json:"name" binding:"required"`
	Venue  string        `json:"venue" binding:"required"`
	Masjid sql.NullInt32 `json:"masjid" binding:"required"`
	Date   time.Time     `json:"date" binding:"required"`
}

func (server *Server) updateEvent(ctx *gin.Context) {
	var req updateEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateEventParams{
		ID:     req.ID,
		Name:   req.Name,
		Venue:  req.Venue,
		Masjid: req.Masjid,
		Date:   req.Date,
	}

	q := db.New(server.db)
	event, err := q.UpdateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type deleteEventRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteEvent(ctx *gin.Context) {
	var req deleteEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q := db.New(server.db)
	err := q.DeleteEvent(ctx, req.ID)
	if err != nil {
		log.Fatal("Error when connecting to databse: ", err)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
