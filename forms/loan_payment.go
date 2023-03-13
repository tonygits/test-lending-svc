package forms

type LoanPayment struct {
	Username   string  `json:"username"`
	LoanAmount float64 `json:"loan_amount"`
}
