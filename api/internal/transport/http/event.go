package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createEventRequest struct {
	Name        string `json:"name" binding:"required`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Prefix      string `json:"prefix" binding:"required`
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var startDate time.Time
	var endDate time.Time
	var err error

	if req.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if req.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

	}

	arg := db.CreateEventParams{
		ID:   uuid.New(),
		Name: req.Name,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		StartDate: sql.NullTime{
			Time: startDate,
		},
		EndDate: sql.NullTime{
			Time: endDate,
		},
		Prefix: req.Prefix,
	}

	event, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type getEventRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reqId, _ := uuid.Parse(req.ID)

	event, err := server.store.GetEvent(ctx, reqId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type listEventRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listEvent(ctx *gin.Context) {
	var req listEventRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEventParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	events, err := server.store.ListEvent(ctx, arg)
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
