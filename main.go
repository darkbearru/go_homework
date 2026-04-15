package main

import (
	"fmt"
	"slices"
	"strings"
)

const rateUSD2EUR = 1.14
const rateUSD2RUB = 69.00
const rateEUR2RUB = rateUSD2RUB / rateUSD2EUR

var currencies = []string{"USD", "EUR", "RUB"}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var curencySrc string = ""
	var currencyDst string = ""
	var amount float64 = 0
	curencySrc = userCurrencyInput("Enter source currency:", curencySrc)
	currencyDst = userCurrencyInput("Enter destination currency:", curencySrc)
	amount = readAmount("Enter amount:")
	fmt.Printf("Amount: %f %s\n", amount, curencySrc)
	fmt.Printf("Converted amount: %f %s\n", calculateAmount(amount, curencySrc, currencyDst), currencyDst)
}

func userCurrencyInput(message string, selected string) string {
	var acceptedCurrencies []string
	var value string
	for _, currency := range currencies {
		if currency != selected {
			acceptedCurrencies = append(acceptedCurrencies, currency)
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
		if slices.Index(acceptedCurrencies, value) == -1 {
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
	switch currencySrc {
	case "USD":
		return calcUSD(amount, currencyDst)
	case "EUR":
		return calcEUR(amount, currencyDst)
	case "RUB":
		return calcRUB(amount, currencyDst)
	}
	return amount
}

func calcUSD(amount float64, currencyDst string) float64 {
	switch currencyDst {
	case "EUR":
		return amount * rateUSD2EUR
	case "RUB":
		return amount * rateUSD2RUB
	}
	return 0
}

func calcEUR(amount float64, currencyDst string) float64 {
	switch currencyDst {
	case "USD":
		return amount / rateUSD2EUR
	case "RUB":
		return amount * rateEUR2RUB
	}
	return 0
}

func calcRUB(amount float64, currencyDst string) float64 {
	switch currencyDst {
	case "USD":
		return amount / rateUSD2RUB
	case "EUR":
		return amount / rateEUR2RUB
	}
	return 0
}
