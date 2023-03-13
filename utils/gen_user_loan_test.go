package utils

import (
	"github.com/stretchr/testify/require"
	"github.com/tonygits/test-lending-svc/entities"
	"testing"
	"time"
)

func TestGenUserLoan(t *testing.T) {
	//loan for product A
	loanProduct := &entities.LoanOffer{
		LoanType: entities.LoadProductTypeA,
		Limit: 1000,
		Interest: 10,
		Tenure: 15,
	}

	userLoan := GenerateUserLoan("user1", loanProduct)
	require.Equal(t, float64(1100), userLoan.LoanBalance)
	require.Equal(t, time.Now().AddDate(0,0, 15).Unix(), userLoan.DueDate.Unix())

	//loan for product B
	loanProduct = &entities.LoanOffer{
		LoanType: entities.LoadProductTypeB,
		Limit: 2500,
		Interest: 12.5,
		Tenure: 30,
	}

	userLoan = GenerateUserLoan("user2", loanProduct)
	require.Equal(t, 2812.5, userLoan.LoanBalance)
	require.Equal(t, time.Now().AddDate(0,0, 30).Unix(), userLoan.DueDate.Unix())
}
