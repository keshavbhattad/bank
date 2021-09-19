package utils

const (
	USD = "USD"
	INR = "INR"
	YEN = "YEN"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, YEN, EUR:
		return true
	}
	return false
}