package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/dhiyaaulauliyaa/learn-go/db/sqlc"
	nullable "github.com/dhiyaaulauliyaa/learn-go/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type createEventRequest struct {
	Name      string    `json:"name" binding:"required"`
	Venue     string    `json:"venue" binding:"required"`
	Community null.Int  `json:"community"`
	Masjid    null.Int  `json:"masjid"`
	Date      time.Time `json:"date" binding:"required"`
}

func eventErrHandling(err error, defaultMsg string) (string, int) {
	if err == sql.ErrNoRows {
		return "Data not found", http.StatusNotFound
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "foreign_key_violation":
			switch pqErr.Constraint {
			case "events_masjid_fkey":
				return "Masjid not found", http.StatusNotFound
			}
		}
	}

	return defaultMsg, http.StatusInternalServerError
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
		return
	}

	arg := db.CreateEventParams{
		Name:   req.Name,
		Venue:  req.Venue,
		Masjid: nullable.NullableToInt32(req.Masjid),
		Date:   req.Date,
	}

	res, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		message, code := eventErrHandling(err, "Create event failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
}

type getEventRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindUri))
		return
	}

	res, err := server.store.GetEvent(ctx, req.ID)
	if err != nil {
		message, code := eventErrHandling(err, "Get event failed")
		ctx.JSON(code, errorResponse(err, message))

		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
}

func (server *Server) getEvents(ctx *gin.Context) {
	res, err := server.store.ListEvents(ctx)
	if err != nil {
		message, code := eventErrHandling(err, "Get events failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(res))
}

type updateEventRequest struct {
	ID     int32     `json:"id" binding:"required"`
	Name   string    `json:"name" binding:"required"`
	Venue  string    `json:"venue" binding:"required"`
	Masjid null.Int  `json:"masjid" binding:"required"`
	Date   time.Time `json:"date" binding:"required"`
}

func (server *Server) updateEvent(ctx *gin.Context) {
	var req updateEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindBody))
		return
	}

	arg := db.UpdateEventParams{
		ID:     req.ID,
		Name:   req.Name,
		Venue:  req.Venue,
		Masjid: nullable.NullableToInt32(req.Masjid),
		Date:   req.Date,
	}

	res, err := server.store.UpdateEvent(ctx, arg)
	if err != nil {
		message, code := eventErrHandling(err, "Update event failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}
	ctx.JSON(http.StatusOK, successResponse(res))
}

type deleteEventRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteEvent(ctx *gin.Context) {
	var req deleteEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, errorBindUri))
		return
	}

	err := server.store.DeleteEvent(ctx, req.ID)
	if err != nil {
		message, code := eventErrHandling(err, "Delete event failed")
		ctx.JSON(code, errorResponse(err, message))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(nil))
}
