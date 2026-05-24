package dto

type CreateDebtRequest struct {
	PersonName  string  `json:"person_name" binding:"required"`
	AccountType string  `json:"account_type" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Note        string  `json:"note" binding:"required"`
}

type UpdateDebtRequest struct {
	PersonName  *string  `json:"person_name" binding:"required"`
	AccountType *string  `json:"account_type" binding:"required"`
	Amount      *float64 `json:"amount" binding:"required"`
	Type        *string  `json:"type" binding:"required"`
	Status      *string  `json:"status" binding:"required"`
	Note        *string  `json:"note" binding:"required"`
}
