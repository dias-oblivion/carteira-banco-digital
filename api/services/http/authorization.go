package services

import (
	"net/http"

	"github.com/dias-oblivion/PicPay-Simplificado/api/utils"
)

type Authorization struct{}

func (Authorization) CheckExternalAuthorization() (string, error) {
	resp, err := http.Get("https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc")
	if err != nil {
		return "", err
	}

	var authorization utils.TransferExternalAuthResponse

	err = utils.GetResponseJson(resp, &authorization)
	if err != nil {
		return "", err
	}

	return authorization.Message, err
}
