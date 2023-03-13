package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tonygits/test-lending-svc/forms"
	"net/http"
)

func (s *Router) loanOfferRequestAPI(ctx *gin.Context) {

	username := ctx.Param("username")

	loanOffers, status, err := userLoanOfferRequest(username)
	if err != nil {
		ctx.JSON(status, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loanOffers)
}

func (s *Router) userOfferAcceptanceAPI(ctx *gin.Context) {

	var form *forms.LoanOfferAcceptance
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userLoan, status, err := userLoanOfferAcceptance(form)
	if err != nil {
		ctx.JSON(status, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userLoan)
}

func (s *Router) userLoanPaymentAPI(ctx *gin.Context) {
	var form *forms.LoanPayment
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, status, err := userLoanPayment(form)
	if err != nil {
		ctx.JSON(status, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Loan repaid successfully",
	})
}
