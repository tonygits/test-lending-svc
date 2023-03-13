package entities

type LoanOffer struct {
	LoanType LoanProductType `json:"loan_type"`
	Limit    float64         `json:"limit"`
	Interest float64         `json:"interest"`
	Tenure   int             `json:"tenure"`
}
