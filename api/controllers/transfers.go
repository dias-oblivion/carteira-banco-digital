package controllers

import (
	"net/http"

	"github.com/dias-oblivion/carteira-banco-digital/api/services"
	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	"github.com/dias-oblivion/carteira-banco-digital/api/utils"
	"github.com/gin-gonic/gin"
)

type Transfer struct{}

func (Transfer) PostTransferBalance(ctx *gin.Context) {
	var req request.TransferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.GetCurrentUserInfo(ctx)

	if err != nil || user.Role != "user" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user permissions"})
		return
	}

	err = services.Transfer{}.TransferBalance(user.ID, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "transfer completed successfully"})

}

func (Transfer) GetTransfersHistoryList(ctx *gin.Context) {
	var req request.TransfersHistoryListRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.GetCurrentUserInfo(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfersHistory, err := services.Transfer{}.GetTransfersHistoryList(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transfersHistory})

}
