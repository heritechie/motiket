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

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
