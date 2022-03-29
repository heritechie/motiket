package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type paymentOrderRequest struct {
	PaymentOptionID int32  `json:"payment_option_id" binding:"required"`
	CustomerOrderID string `json:"customer_order_id" binding:"required"`
}

func (server *Server) paymentOrder(ctx *gin.Context) {
	var req paymentOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	coUUID, _ := uuid.Parse(req.CustomerOrderID)

	arg := db.PaymentTxParams{
		CustomerOrderID: coUUID,
		PaymentOptionID: req.PaymentOptionID,
	}

	event, err := server.store.PaymentTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type paymentOrderConfirmationRequest struct {
	PaymentID    string `json:"payment_id" binding:"required"`
	Status       string `json:"status" binding:"required,oneof=SUCCESS FAILED"`
	FailedReason string `json:"failed_reason"`
}

func (server *Server) paymentOrderConfirmation(ctx *gin.Context) {
	var req paymentOrderConfirmationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	paymentUUID, _ := uuid.Parse(req.PaymentID)

	arg := db.PaymentConfirmationTxParams{
		PaymentID:    paymentUUID,
		Status:       req.Status,
		FailedReason: req.FailedReason,
	}

	event, err := server.store.PaymentConfirmationTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}
