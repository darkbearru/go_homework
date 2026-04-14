package main

import "fmt"

const rateUSD2EUR = 1.14
const rateUSD2RUB = 69.00
const rateEUR2RUB = rateUSD2RUB / rateUSD2EUR

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var amountEUR float64 = readAmount()
	amount := amountEUR * rateEUR2RUB
	amount = calculateAmount(amount, "EUR", "RUB")
	println(amount)
}

func readAmount() float64 {
	var amount float64
	fmt.Scan(&amount)
	return amount
}

func calculateAmount(amount float64, currencySrc string, currencyDst string) float64 {
	return amount
}
