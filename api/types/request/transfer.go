package types

type PaginationListRequest struct {
	Limit  int64 `json:"limit" binding:"required"`
	Offset int64 `json:"offset" binding:"required"`
}

type TransferRequest struct {
	PayeeID     int64   `json:"payee_id" binding:"required"`
	Value       float64 `json:"value" binding:"required"`
	Description string  `json:"description"`
}

type TransfersHistoryListRequest struct {
	PaginationListRequest
}
