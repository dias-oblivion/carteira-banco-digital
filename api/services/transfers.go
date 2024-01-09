package services

import (
	"errors"
	"fmt"

	services "github.com/dias-oblivion/carteira-banco-digital/api/services/http"
	request "github.com/dias-oblivion/carteira-banco-digital/api/types/request"
	types "github.com/dias-oblivion/carteira-banco-digital/api/types/response"
	"github.com/dias-oblivion/carteira-banco-digital/database/repositories"
	"github.com/jinzhu/copier"
)

type Transfer struct{}

func (Transfer) TransferBalance(userID int64, transfer request.TransferRequest) error {

	authorization, err := services.Authorization{}.CheckExternalAuthorization()

	if err != nil || authorization != "Autorizado" {
		return errors.New("error: unauthorized transfer")
	}

	err = repositories.Transfer{}.TransferBalance(userID, transfer)
	if err != nil {
		return err
	}

	_, err = services.Notification{}.SendNotification()

	if err != nil {
		// TODO ~ implementar sistema de fila de notificações
		fmt.Println("error: can't send notification")
	}

	return nil
}

func (Transfer) GetTransfersHistoryList(userID int64) ([]types.TransferHistoryResponse, error) {
	transfersHistory, err := repositories.Transfer{}.GetTransfersHistory(userID)
	if err != nil {
		return []types.TransferHistoryResponse{}, err
	}

	var transferHistoryResponse []types.TransferHistoryResponse
	copier.Copy(&transferHistoryResponse, &transfersHistory)

	return transferHistoryResponse, nil
}
