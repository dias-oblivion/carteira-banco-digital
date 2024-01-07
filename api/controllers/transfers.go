package controllers

import (
	"net/http"

	"github.com/dias-oblivion/PicPay-Simplificado/api/services"
	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/api/utils"
	"github.com/gin-gonic/gin"
)

type Transfer struct{}

func (Transfer) TransferBalance(ctx *gin.Context) {
	var req request.Transfer

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetCurrentUserID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.Transfer{}.TransferBalance(userId, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "transfer completed successfully"})

}
