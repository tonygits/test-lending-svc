package utils

import (
	"github.com/tonygits/test-lending-svc/entities"
	"time"
)

func GenerateUserLoan(username string, loanProduct *entities.LoanOffer) *entities.UserLoan {
	userLoan := &entities.UserLoan{
		Username:   username,
	}

	dueDate := time.Now().AddDate(0, 0, loanProduct.Tenure)
	userLoan.DueDate = dueDate

	loanAmount := loanProduct.Limit*(1+(loanProduct.Interest/100))
	userLoan.LoanBalance = loanAmount

	return userLoan
}

