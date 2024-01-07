package services

import (
	"io"
	"net/http"
)

type Notification struct{}

func (Notification) SendNotification() (string, error) {
	resp, err := http.Get("https://run.mocky.io/v3/54dc2cf1-3add-45b5-b5a9-6bf7e7f1f4a6")
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
