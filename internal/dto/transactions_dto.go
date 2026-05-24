package dto

type CreateTransactionRequest struct {
	AccountID   int     `json:"account_id"`
	AccountType string  `json:"account_type" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Note        string  `json:"note"`
}

type FilterTransaction struct {
	Type        string `json:"type"`
	Category    string `json:"category"`
	AccountType string `json:"account_type"`
}
