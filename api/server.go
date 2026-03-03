package api

import (
	"my-bank-app/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *sqlc.Store
	router *gin.Engine
}
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func NewServer(store *sqlc.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
