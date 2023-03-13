package utils

import "github.com/tonygits/test-lending-svc/entities"

func GenerateLoanOffers(userLimit float64) []*entities.LoanOffer {

	loanOffer1 := &entities.LoanOffer{
		LoanType: entities.LoadProductTypeA,
		Limit:    1000,
		Interest: 10,
		Tenure:   15,
	}

	loanOffer2 := &entities.LoanOffer{
		LoanType: entities.LoadProductTypeB,
		Limit:    2500,
		Interest: 12.5,
		Tenure:   30,
	}

	if userLimit >= 1000 && userLimit < 2500 {
		return []*entities.LoanOffer{loanOffer1}
	}

	if userLimit >= 2500 {
		return []*entities.LoanOffer{loanOffer1, loanOffer2}
	}

	return []*entities.LoanOffer{}
}
