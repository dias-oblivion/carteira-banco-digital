package controllers

import (
	"net/http"

	"github.com/dias-oblivion/carteira-banco-digital/api/services"
	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (User) PostCreateUser(ctx *gin.Context) {
	var user request.CreateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := services.User{}.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": userId})
}

func (User) PostLogin(ctx *gin.Context) {
	var credentials request.Login
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.User{}.Login(credentials)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
