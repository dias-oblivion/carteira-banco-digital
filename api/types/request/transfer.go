package types

type Transfer struct {
	PayeeID     int64   `json:"payee_id" binding:"required"`
	Value       float64 `json:"value" binding:"required"`
	Description string  `json:"description"`
}
