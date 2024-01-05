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

	// Public Routes
	public.POST("/login", user.Login)

	// Private Routes
	private.POST("/user", user.CreateUser)

	router.Run(s.port)
}
