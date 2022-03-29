package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createTicketRequest struct {
	SerialNumber     string `json:"serial_number"`
	TicketCategoryID string `json:"ticket_category_id" binding:"required`
	EventID          string `json:"event_id" binding:"required`
}

func (server *Server) createTicket(ctx *gin.Context) {
	var req createTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tcUUID, _ := uuid.Parse(req.TicketCategoryID)
	eventIDUUID, _ := uuid.Parse(req.EventID)

	arg := db.CreateTicketParams{
		SerialNumber:     req.SerialNumber,
		TicketCategoryID: tcUUID,
		EventID:          eventIDUUID,
	}

	event, err := server.store.CreateTicket(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type getTicketRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getTicket(ctx *gin.Context) {
	var req getTicketRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reqId, _ := uuid.Parse(req.ID)

	ticket, err := server.store.GetTicket(ctx, reqId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ticket)
}

type listTicketRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTicket(ctx *gin.Context) {
	var req listTicketRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTicketParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listTicket, err := server.store.ListTicket(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listTicket)
}
