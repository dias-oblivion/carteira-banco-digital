package api

import (
	"github.com/dias-oblivion/PicPay-Simplificado/api/routes"
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

	router.POST("/user", routes.CreateUser)

	// router.GET("/transactions/history")

	router.Run(s.port)
}
