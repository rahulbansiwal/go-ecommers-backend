package api

import (
	"ecom/db/sqlc"
	"ecom/db/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  sqlc.Store
	router *gin.Engine
}

func NewServer(config util.Config, store sqlc.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRoutes()

	return server, nil
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRoutes() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.POST("/user", s.CreateUser)
	router.GET("/user/:id",s.GetUser)
	router.POST("/user/:username",s.UpdateUser)
	// ADD NEW ROUTES

	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
