package handler

import "github.com/gin-gonic/gin"

func CreateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ok!",
	})
}
