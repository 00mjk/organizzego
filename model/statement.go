package model

type Statement struct {
	Description   string `json:"description"`
	Date          string `json:"date"`
	AmountInCents int    `json:"amount_cents"`
}
