package model

type Statement struct {
	Description            string                  `json:"description"`
	Date                   string                  `json:"date"`
	AmountInCents          int                     `json:"amount_cents"`
	Paid                   bool                    `json:"paid"`
	AccountID              int                     `json:"account_id"`
	CreditCardID           int                     `json:"credit_card_id"`
	CategoryID             int                     `json:"category_id"`
	RecurrenceAttributes   *RecurrenceAttributes   `json:"recurrence_attributes"`
	InstallmentsAttributes *InstallmentsAttributes `json:"installments_attributes"`
}

func (statement Statement) New(description string, date string, amountInCents int, paid bool) Statement {
	statement.Description = description
	statement.Date = date
	statement.AmountInCents = amountInCents
	statement.Paid = paid
	return statement
}
