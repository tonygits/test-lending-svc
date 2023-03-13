package services

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

// NewRouter The function creates a new http server and set up routing
func NewRouter() (*Router, error) {
	router := &Router{}
	router.setRouter()
	return router, nil
}

func (s *Router) setRouter() {
	router := gin.Default()
	//Loan Request
	router.GET("/users/:username", s.loanOfferRequestAPI)
	router.POST("/loan_offer_acceptance", s.userOfferAcceptanceAPI)

	//Loan payment
	router.POST("/loan_payment", s.userLoanPaymentAPI)

	//add routes to router
	s.router = router
}

// Start The function runs the http server on a specific address
func (s *Router) Start(address string) error {
	return s.router.Run(address)
}

// handle the error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
