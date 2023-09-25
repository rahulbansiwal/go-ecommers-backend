package api

import (
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  sqlc.Store
	router *gin.Engine
	paseto *token.PasetoMaker
}

func NewServer(config util.Config, store sqlc.Store) (*Server, error) {
	paseto, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: config,
		store:  store,
		paseto: paseto,
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
	router.POST("/user/login", s.LoginUser)

	authRoutes := router.Group("/", AuthMiddleware(s.paseto))
	authRoutes.POST("/user/:username", s.UpdateUser)
	authRoutes.GET("/user/:id", s.GetUser)
	authRoutes.GET("/logout", s.LogoutUser)
	authRoutes.GET("/logout/all", s.LogoutUserFromAllDevice)
	// ADD NEW ROUTES

	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// TODO: ADD GITHUB ACTIONS
