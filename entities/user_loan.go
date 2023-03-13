package entities

import "time"

type UserLoan struct {
	Username    string    `json:"username"`
	LoanBalance float64   `json:"loan_balance"`
	DueDate     time.Time `json:"due_date"`
}
