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

	user := controllers.User{}

	router.POST("/user", user.CreateUser)

	// router.GET("/transactions/history")

	router.Run(s.port)
}
