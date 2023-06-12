package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenLoanOffers(t *testing.T) {
	userLimit1 := 3000.00
	userLimit2 := 1500.00
	userLimit3 := 500.00
	userLimit4 := 1000.00
	userLimit5 := 2500.00

	//userLimit above product B
	loanOffers := GenerateLoanOffers(userLimit1)
	require.Equal(t, 2, len(loanOffers))

	//userLimit above product A, but below product B
	loanOffers = GenerateLoanOffers(userLimit2)
	require.Equal(t, 1, len(loanOffers))

	//userLimit below product A
	loanOffers = GenerateLoanOffers(userLimit3)
	require.Equal(t, 0, len(loanOffers))

	//userLimit equal product A
	loanOffers = GenerateLoanOffers(userLimit4)
	require.Equal(t, 1, len(loanOffers))

	//user limit equal product B
	loanOffers = GenerateLoanOffers(userLimit5)
	require.Equal(t, 2, len(loanOffers))
}
