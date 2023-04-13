package handlers

import "fmt"

func CurrencySim(in string) (string, error) {
	switch in {
	case "BRL":
		return "R$", nil
	case "USD":
		return "$", nil
	case "EUR":
		return "€", nil
	case "BTC":
		return "₿", nil
	default:
		return "", fmt.Errorf("Currency not supported: %s", in)
	}
}