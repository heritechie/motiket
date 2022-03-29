package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createTicketCategoryRequest struct {
	Name      string `json:"name" binding:"required"`
	Area      string `json:"area"`
	Qty       int64  `json:"qty" binding:"required"`
	Price     int64  `json:"price" binding:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Prefix    string `json:"prefix" binding:"required"`
	EventID   string `json:"event_id" binding:"required"`
}

func (server *Server) createTicketCategory(ctx *gin.Context) {
	var req createTicketCategoryRequest
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

	eventIDUUID, _ := uuid.Parse(req.EventID)

	arg := db.CreateTicketCategoryParams{
		ID:    uuid.New(),
		Name:  req.Name,
		Qty:   req.Qty,
		Price: req.Price,
		StartDate: sql.NullTime{
			Time: startDate,
		},
		EndDate: sql.NullTime{
			Time: endDate,
		},
		Prefix: req.Prefix,
		Area: sql.NullString{
			String: req.Area,
			Valid:  true,
		},
		EventID: eventIDUUID,
	}

	event, err := server.store.CreateTicketCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type getTicketCategoryRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getTicketCategory(ctx *gin.Context) {
	var req getTicketCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reqId, _ := uuid.Parse(req.ID)

	ticketCategory, err := server.store.GetTicketCategory(ctx, reqId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ticketCategory)
}

type listTicketCategoryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTicketCategory(ctx *gin.Context) {
	var req listEventRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTicketCategoryParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listTicketCategory, err := server.store.ListTicketCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listTicketCategory)
}
