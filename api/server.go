package api

import (
	"ecom/db/sqlc"
	"ecom/db/util"
	"ecom/token"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      sqlc.Store
	router     *gin.Engine
	paseto     *token.PasetoMaker
	s3uploader *manager.Uploader
}

func NewServer(config util.Config, store sqlc.Store, uploader *manager.Uploader) (*Server, error) {
	paseto, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		store:      store,
		paseto:     paseto,
		s3uploader: uploader,
	}
	server.setupRoutes()

	return server, nil
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRoutes() {
	router := gin.New()
	router.Use(gin.Recovery())
	f, err := loggingInFile(s.config)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(f)
	}
	router.Use(jsonLoggerMiddleware())
	router.MaxMultipartMemory = 8 << 20
	gin.SetMode(gin.DebugMode)
	router.POST("/user", s.CreateUser)
	router.POST("/user/login", s.LoginUser)
	router.POST("/refreshtoken", s.renewAccessToken)

	authRoutes := router.Group("/", AuthMiddleware(s.paseto))
	authRoutes.POST("/address", s.addAddress)
	authRoutes.GET("/address/:id", s.GetAddress)
	authRoutes.DELETE("/address/:id", s.deleteAddress)
	authRoutes.POST("/user/:username", s.UpdateUser)
	authRoutes.GET("/user/:id", s.GetUser)
	authRoutes.GET("/logout", s.LogoutUser)
	authRoutes.GET("/logout/all", s.LogoutUserFromAllDevice)
	authRoutes.POST("/item", s.createItem)

	// ADD NEW ROUTES

	s.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
