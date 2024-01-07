package api

import (
	"github.com/dias-oblivion/PicPay-Simplificado/api/controllers"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	port string
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{
		port: port,
	}
}

func (s *APIServer) Start() {
	router := gin.Default()

	public := router.Group("/api/v1")
	private := router.Group("/api/v1").Use(AuthMiddleware())

	user := controllers.User{}
	transfer := controllers.Transfer{}

	// Public Routes
	public.POST("/login", user.Login)
	public.POST("/user", user.CreateUser)

	// Private Routes
	private.POST("/transfer", transfer.TransferBalance)

	router.Run(s.port)
}
