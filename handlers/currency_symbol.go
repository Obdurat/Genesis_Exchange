package handlers

import "fmt"

func CurrencySim(in string) (string, error) {
	var simbol string
	var err error
	switch in {
	case "BRL":
		simbol = "R$"; break
	case "USD":
		simbol = "$"; break
	case "EUR":
		simbol = "€"; break
	case "BTC":
		simbol = "₿"; break
	default:
		err = fmt.Errorf("Currency not supported: %s", in)
	}
	return simbol, err
}