package model

type CreditCard struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	CardNetwork string      `json:"card_network"`
	ClosingDay  int         `json:"closing_day"`
	DueDay      int         `json:"due_day"`
	LimitCents  int         `json:"limit_cents"`
	Kind        string      `json:"kind"`
	Archived    bool        `json:"archived"`
	Default     bool        `json:"default"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}
