package types

type TransferHistoryResponse struct {
	ID           int64   `json:"id"`
	UserName     string  `json:"user_name"`
	UserDocument string  `json:"user_document"`
	Description  string  `json:"description"`
	Value        float64 `json:"value"`
	Type         string  `json:"type"`
}
