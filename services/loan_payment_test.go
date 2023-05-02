package services

import (
	"github.com/stretchr/testify/require"
	"github.com/tonygits/test-lending-svc/forms"
	"net/http"
	"testing"
)

func TestLoanPayment(t *testing.T) {

	//loan payment with correct details
	form := &forms.LoanPayment{
		Username: "user1",
		LoanAmount: 1100,
	}

	userLoan, status, err := userLoanPayment(form)
	require.NoError(t, err)
	require.Equal(t, float64(1100), userLoan.LoanBalance)
	require.Equal(t, http.StatusOK, status)

	//loan payment with wrong loan amount
	form = &forms.LoanPayment{
		Username: "user1",
		LoanAmount: 1500,
	}

	_, status, err = userLoanPayment(form)
	require.Error(t, err)
	require.Equal(t, "exact loan amount must be paid in full", err.Error())
	require.Equal(t, http.StatusBadRequest, status)

	//wrong username
	form = &forms.LoanPayment{
		Username: "user15",
		LoanAmount: 1100,
	}

	_, status, err = userLoanPayment(form)
	require.Error(t, err)
	require.Equal(t, "user does not exist", err.Error())
	require.Equal(t, http.StatusNotFound, status)
}
