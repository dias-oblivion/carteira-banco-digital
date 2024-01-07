package services

import (
	"errors"
	"fmt"

	services "github.com/dias-oblivion/PicPay-Simplificado/api/services/http"
	request "github.com/dias-oblivion/PicPay-Simplificado/api/types/request"
	"github.com/dias-oblivion/PicPay-Simplificado/database/repositories"
)

type Transfer struct{}

func (Transfer) TransferBalance(userID int64, transfer request.Transfer) error {

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
