package utils

type TransferExternalAuthResponse struct {
	Message string `json:"message"`
}

type UserInfo struct {
	ID    int64
	Email string
	Role  string
}
