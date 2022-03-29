package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createCustomerOrderRequest struct {
	CustomerID        string `json:"customer_id" binding:"required"`
	CustomerPaymentID string `json:"customer_payment_id" binding:"required"`
	TotalPrice        int64  `json:"total_price"`
	Discount          int32  `json:"discount"`
	FinalPrice        int64  `json:"final_price"`
}

func (server *Server) createCustomerOrder(ctx *gin.Context) {
	var req createCustomerOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cUUID, _ := uuid.Parse(req.CustomerID)
	cpUUID, _ := uuid.Parse(req.CustomerPaymentID)

	arg := db.CreateCustomerOrderParams{
		CustomerID:        cUUID,
		CustomerPaymentID: cpUUID,
		TotalPrice:        req.TotalPrice,
		Discount: sql.NullInt32{
			Int32: req.Discount,
			Valid: true,
		},
		FinalPrice: req.FinalPrice,
	}

	event, err := server.store.CreateCustomerOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}

type getCustomerOrderRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getCustomerOrder(ctx *gin.Context) {
	var req getCustomerOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reqId, _ := uuid.Parse(req.ID)

	customerOrder, err := server.store.GetCustomerOrder(ctx, reqId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customerOrder)
}

type listCustomerOrderRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCustomerOrder(ctx *gin.Context) {
	var req listCustomerOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCustomerOrderParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listCustomerOrder, err := server.store.ListCustomerOrder(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listCustomerOrder)
}
