package utils

func CreditUserWallet(walletBalance float64) string {
	if walletBalance == 5000 {
		return "failed"
	}
	return "success"
}
