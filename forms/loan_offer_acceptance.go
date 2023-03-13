package forms

import "github.com/tonygits/test-lending-svc/entities"

type LoanOfferAcceptance struct {
	Username string                   `json:"username"`
	LoanType entities.LoanProductType `json:"loan_type"`
}
