package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type orderTicketParams struct {
	Qty      int32  `json:"qty"`
	TicketID string `json:"ticket_id"`
}

type checkoutOrderRequest struct {
	CustomerID      string               `json:"customer_id" binding:"required"`
	TotalPrice      int64                `json:"total_price" binding:"required"`
	Discount        int32                `json:"discount"`
	FinalPrice      int64                `json:"final_price"`
	ListOrderTicket []*orderTicketParams `json:"list_order_ticket" binding:"required"`
}

func (server *Server) checkoutOrder(ctx *gin.Context) {
	var req checkoutOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	custUUID, _ := uuid.Parse(req.CustomerID)

	var listOrderTicket []db.OrderTicketParams

	for _, ot := range req.ListOrderTicket {
		ticketUUID, _ := uuid.Parse(ot.TicketID)
		listOrderTicket = append(listOrderTicket, db.OrderTicketParams{
			Qty:      ot.Qty,
			TicketID: ticketUUID,
		})
	}

	arg := db.CheckoutTxParams{
		CustomerID:      custUUID,
		TotalPrice:      req.FinalPrice,
		Discount:        req.Discount,
		FinalPrice:      req.FinalPrice,
		ListOrderTicket: listOrderTicket,
	}

	event, err := server.store.CheckoutTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type getOrderByCustomerOrderIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getOrderByCustomerOrderId(ctx *gin.Context) {
	var req getOrderByCustomerOrderIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, _ := uuid.Parse(req.ID)

	order, err := server.store.GetCustomerOrder(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
