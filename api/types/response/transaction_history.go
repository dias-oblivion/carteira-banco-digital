package types

type TransactionHistory struct {
	UserName     string  `json:"user_name"`
	UserDocument string  `json:"document_number"`
	Description  string  `json:"description"`
	Value        float32 `json:"value"`
	Type         string  `json:"type"`
}
