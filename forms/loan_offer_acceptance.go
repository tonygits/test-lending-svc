package forms

import "github.com/tonygits/test-lending-svc/entities"

type LoanOfferAcceptance struct {
	Username string                   `binding:"required" json:"username"`
	LoanType entities.LoanProductType `binding:"required" json:"loan_type"`
}