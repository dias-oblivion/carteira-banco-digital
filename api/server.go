package api

import (
	"github.com/dias-oblivion/carteira-banco-digital/api/controllers"
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
	public.POST("/login", user.PostLogin)
	public.POST("/user", user.PostCreateUser)

	// Private Routes
	private.POST("/transfer", transfer.PostTransferBalance)
	private.GET("/transfers", transfer.GetTransfersHistoryList)

	router.Run(s.port)
}
