package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/procyon-projects/chrono"
	"github.com/tonygits/test-lending-svc/entities"
	"github.com/tonygits/test-lending-svc/forms"
	"github.com/tonygits/test-lending-svc/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

func userLoanOfferRequest(username string) ([]*entities.LoanOffer, int, error) {
	if len(username) == 0 {
		err := errors.New("specify the username for the user")
		return nil, http.StatusBadRequest, err
	}

	loanOffers, status, err := getUserLoanProfile(username)
	if err != nil {
		return loanOffers, status, err
	}

	return loanOffers, http.StatusOK, err
}

func userLoanOfferAcceptance(form *forms.LoanOfferAcceptance) (*entities.UserLoan, int, error) {
	loanOffers, status, err := getUserLoanProfile(form.Username)
	if err != nil {
		return nil, status, err
	}

	if !form.LoanType.IsValid() {
		err := errors.New("loan product is invalid")
		return nil, http.StatusBadRequest, err
	}

	loanProduct, isProductIncluded := loanOffersContain(loanOffers, form.LoanType)
	if !isProductIncluded {
		err := errors.New("loan product is invalid")
		return nil, http.StatusBadRequest, err
	}

	userLoan := utils.GenerateUserLoan(form.Username, loanProduct)
	if userLoan.DueDate.Before(time.Now()) {
		err := errors.New("due date for the loan is invalid")
		return nil, http.StatusBadRequest, err
	}

	log.Printf("%s wallet credited with %f", form.Username, loanProduct.Limit)

	//send notification
	message := fmt.Sprintf("Loan offer selection was successful. Your wallet has been credited with %f", loanProduct.Limit)
	utils.SendNotification(message)

	//schedule automated payment deduction from wallet
	scheduleLoanDeductionFromWallet(userLoan, userLoan.DueDate)
	return userLoan, http.StatusOK, nil
}

func getUserLoanProfile(username string) ([]*entities.LoanOffer, int, error) {
	//check if user exists
	userMap := getUserAccounts()
	if _, exists := userMap[username]; !exists {
		err := errors.New("user does not exist")
		return nil, http.StatusNotFound, err
	}

	//check if user is qualified for loan offers
	loanOffers := utils.GenerateLoanOffers(userMap[username].UserLimit)
	if len(loanOffers) == 0 {
		message := "We have no loan offers for you today"
		utils.SendNotification(message)
		err := errors.New("we have no loan offers for you today")
		return nil, http.StatusNotFound, err
	}

	return loanOffers, http.StatusOK, nil
}

func getUserAccounts() map[string]*entities.User {
	userMap := make(map[string]*entities.User)
	user1 := utils.GenerateUser("user1", 1000)
	user2 := utils.GenerateUser("user2", 2500)
	user3 := utils.GenerateUser("user3", 2000)
	user4 := utils.GenerateUser("user4", 4000)
	user5 := utils.GenerateUser("user5", 1500)
	user6 := utils.GenerateUser("user6", 500)
	user7 := utils.GenerateUser("user7", 8000)
	userMap["user1"] = user1
	userMap["user2"] = user2
	userMap["user3"] = user3
	userMap["user4"] = user4
	userMap["user5"] = user5
	userMap["user6"] = user6
	userMap["user7"] = user7
	return userMap
}

func loanOffersContain(loanOffers []*entities.LoanOffer, loanSelected entities.LoanProductType) (*entities.LoanOffer, bool) {
	if len(loanOffers) == 0 {
		return nil, false
	}

	for _, v := range loanOffers {
		if strings.ToLower(v.LoanType.String()) == strings.ToLower(loanSelected.String()) {
			return v, true
		}
	}

	return nil, false
}

func scheduleLoanDeductionFromWallet(userLoan *entities.UserLoan, dueDate time.Time) {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	_, _ = taskScheduler.Schedule(func(ctx context.Context) {
		log.Printf("You loan of %f has been deducted automatically from wallet after expiry of period", userLoan.LoanBalance)
	}, chrono.WithTime(dueDate.Add(time.Hour*24)))
}
