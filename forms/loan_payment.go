package forms

type LoanPayment struct {
	Username   string  `binding:"required" json:"username"`
	LoanAmount float64 `binding:"required" json:"loan_amount"`
}
