package router

import (
	"github.com/dias-oblivion/PicPay-Simplificado/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user", handler.CreateUserHandler)
	}
}
