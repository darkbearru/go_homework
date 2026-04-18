package main

import (
	"fmt"
	"strings"
)

const rateUSD2EUR = 1.14
const rateUSD2RUB = 69.00
const rateEUR2RUB = rateUSD2RUB / rateUSD2EUR

var ratesMap = map[string]float64{
	"USD2RUB": rateUSD2RUB,
	"USD2EUR": rateUSD2EUR,
	"EUR2RUB": rateEUR2RUB,
	"EUR2USD": 1 / rateUSD2EUR,
	"RUB2USD": 1 / rateUSD2RUB,
	"RUB2EUR": 1 / rateEUR2RUB,
}

type CurrencyType = map[string]int

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var currencies = CurrencyType{"USD": 1, "EUR": 1, "RUB": 1}

	var currencySrc string = ""
	var currencyDst string = ""
	var amount float64 = 0
	currencySrc = userCurrencyInput("Enter source currency:", &currencySrc, &currencies)
	currencyDst = userCurrencyInput("Enter destination currency:", &currencySrc, &currencies)
	amount = readAmount("Enter amount:")
	fmt.Printf("Amount: %.2f %s\n", amount, currencySrc)
	fmt.Printf("Converted amount: %.2f %s\n", calculateAmount(amount, currencySrc, currencyDst), currencyDst)
}

func userCurrencyInput(message string, selected *string, currencies *CurrencyType) string {
	var value string
	var acceptedCurrencies []string
	for value := range *currencies {
		if value != *selected {
			acceptedCurrencies = append(acceptedCurrencies, value)
		}
	}
	for {
		fmt.Printf(message+"(%s)\n", strings.Join(acceptedCurrencies, ", "))
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("...Invalid currency")
			continue
		}
		value = strings.ToUpper(value)
		_, ok := (*currencies)[value]
		if !ok {
			fmt.Println("...Invalid currency")
			continue
		}
		break
	}
	return value
}

func readAmount(message string) float64 {
	var amount float64
	fmt.Println(message)
	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("...Invalid amount")
			continue
		}
		break
	}
	return amount
}

func calculateAmount(amount float64, currencySrc string, currencyDst string) float64 {
	fmt.Println(currencySrc + "2" + currencyDst)
	return amount * ratesMap[currencySrc+"2"+currencyDst]
}
