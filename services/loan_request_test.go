package services

import (
	"github.com/stretchr/testify/require"
	"github.com/tonygits/test-lending-svc/entities"
	"github.com/tonygits/test-lending-svc/forms"
	"net/http"
	"testing"
)

func TestUserLoanOfferRequest(t *testing.T) {
	users := getUserAccounts()
	require.Equal(t, 7, len(users))

	//user limit between 1000 and 2500
	loanOffers, status, err := userLoanOfferRequest("user5")
	require.NoError(t, err)
	require.Equal(t, 1, len(loanOffers))
	require.Equal(t, "product_a", loanOffers[0].LoanType.String())
	require.Equal(t, http.StatusOK, status)

	//user limit below 1000
	_, status, err = userLoanOfferRequest("user6")
	require.Error(t, err)
	require.Equal(t, "we have no loan offers for you today", err.Error())
	require.Equal(t, http.StatusNotFound, status)

	//user limit equal 2500
	loanOffers, status, err = userLoanOfferRequest("user2")
	require.NoError(t, err)
	require.Equal(t, 2, len(loanOffers))
	require.Equal(t, http.StatusOK, status)

	//user limit above 2500
	loanOffers, status, err = userLoanOfferRequest("user7")
	require.NoError(t, err)
	require.Equal(t, 2, len(loanOffers))
	require.Equal(t, http.StatusOK, status)

	//user does not exist
	loanOffers, status, err = userLoanOfferRequest("user10")
	require.Error(t, err, "user does not exist")
	require.Equal(t, 0, len(loanOffers))
	require.Equal(t, http.StatusNotFound, status)

	//username is empty
	loanOffers, status, err = userLoanOfferRequest("")
	require.Error(t, err)
	require.Equal(t, 0, len(loanOffers))
	require.Equal(t, http.StatusBadRequest, status)
}

func TestUserLoanOfferAcceptance(t *testing.T) {

	//user1 can accept loan product A
	form := &forms.LoanOfferAcceptance{
		Username: "user1",
		LoanType: entities.LoadProductTypeA,
	}
	userLoan, status, err := userLoanOfferAcceptance(form)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, float64(1100), userLoan.LoanBalance)

	//user1 does not qualify for loan product B
	form = &forms.LoanOfferAcceptance{
		Username: "user1",
		LoanType: entities.LoadProductTypeB,
	}
	_, _, err = userLoanOfferAcceptance(form)
	require.Error(t, err)
	require.Equal(t, "loan product is invalid", err.Error())

	//user2 qualifies for loan product B and can accept loan offer
	form = &forms.LoanOfferAcceptance{
		Username: "user2",
		LoanType: entities.LoadProductTypeB,
	}
	userLoan, status, err = userLoanOfferAcceptance(form)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, 2812.5, userLoan.LoanBalance)

	//user2 qualifies for loan product A and can accept loan offer
	form = &forms.LoanOfferAcceptance{
		Username: "user2",
		LoanType: entities.LoadProductTypeA,
	}
	userLoan, status, err = userLoanOfferAcceptance(form)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, float64(1100), userLoan.LoanBalance)
}
