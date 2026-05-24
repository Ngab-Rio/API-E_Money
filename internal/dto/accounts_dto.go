package dto

type GetAccountByUserIDResponse struct {
	UserID  int     `json:"user_id"`
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
}

type GetSummaryResponse struct {
	UserID       int     `json:"user_id"`
	TotalBalance float64 `json:"total_balance"`
}
