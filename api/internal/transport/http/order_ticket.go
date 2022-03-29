package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createOrderTicketRequest struct {
	Qty             int32  `json:"qty"`
	TicketID        string `json:"ticket_id"`
	CustomerOrderID string `json:"customer_order_id"`
}

func (server *Server) createOrderTicket(ctx *gin.Context) {
	var req createOrderTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ticketUUID, _ := uuid.Parse(req.TicketID)
	coUUID, _ := uuid.Parse(req.CustomerOrderID)

	arg := db.CreateOrderTicketParams{
		Qty:             req.Qty,
		TicketID:        ticketUUID,
		CustomerOrderID: coUUID,
	}

	orderTicket, err := server.store.CreateOrderTicket(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, orderTicket)
}

type listOrderTicketRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listOrderTicket(ctx *gin.Context) {
	var req listOrderTicketRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListOrderTicketParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listOrderTicket, err := server.store.ListOrderTicket(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listOrderTicket)
}
