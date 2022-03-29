package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type createPaymentOptionRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createPaymentOption(ctx *gin.Context) {
	var req createPaymentOptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	paymentOption, err := server.store.CreatePaymentOption(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, paymentOption)
}

type getPaymentOptionRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getPaymentOption(ctx *gin.Context) {
	var req getPaymentOptionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	paymentOption, err := server.store.GetPaymentOption(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, paymentOption)
}

type listPaymentOptionRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPaymentOption(ctx *gin.Context) {
	var req listPaymentOptionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPaymentOptionParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listPaymentOption, err := server.store.ListPaymentOption(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listPaymentOption)
}
