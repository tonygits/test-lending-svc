package services

import (
	"errors"
	"github.com/tonygits/test-lending-svc/entities"
	"github.com/tonygits/test-lending-svc/forms"
	"github.com/tonygits/test-lending-svc/utils"
	"net/http"
	"time"
)

func userLoanPayment(form *forms.LoanPayment) (*entities.UserLoan, int, error) {
	loanOffers, status, err := getUserLoanProfile(form.Username)
	if err != nil {
		return nil, status, err
	}

	if form.LoanAmount != 1100 && form.LoanAmount != 2812.5 {
		err := errors.New("exact loan amount must be paid in full")
		return nil, http.StatusBadRequest, err
	}

	userLoan := utils.GenerateUserLoan(form.Username, loanOffers[0])
	if userLoan.DueDate.Before(time.Now()) {
		err := errors.New("due date for the loan is invalid")
		return userLoan, http.StatusBadRequest, err
	}

	return userLoan, http.StatusOK, nil
}
