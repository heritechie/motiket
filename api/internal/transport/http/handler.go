package http

import (
	"github.com/gin-gonic/gin"
	db "github.com/heritechie/motiket/api/internal/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	apiV1 := "/api/v1/"

	router.POST(apiV1+"events", server.createEvent)
	router.GET(apiV1+"events/:id", server.getEvent)
	router.GET(apiV1+"events", server.listEvent)

	router.POST(apiV1+"ticket-categories", server.createTicketCategory)
	router.GET(apiV1+"ticket-categories/:id", server.getTicketCategory)
	router.GET(apiV1+"ticket-categories", server.listTicketCategory)

	router.POST(apiV1+"tickets", server.createTicket)
	router.GET(apiV1+"tickets/:id", server.getTicket)
	router.GET(apiV1+"tickets", server.listTicket)

	router.POST(apiV1+"customers", server.createCustomer)
	router.GET(apiV1+"customers/:id", server.getCustomer)
	router.GET(apiV1+"customers", server.listCustomer)

	router.POST(apiV1+"payment-options", server.createPaymentOption)
	router.GET(apiV1+"payment-options/:id", server.getPaymentOption)
	router.GET(apiV1+"payment-options", server.listPaymentOption)

	router.POST(apiV1+"checkout-order", server.checkoutOrder)
	router.GET(apiV1+"orders/:id", server.getOrderByCustomerOrderId)

	router.POST(apiV1+"payment-order", server.paymentOrder)
	router.POST(apiV1+"payment-order/confirmation", server.paymentOrderConfirmation)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
