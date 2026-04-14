package main

const rateUSD2EUR = 1.14
const rateUSD2RUB = 69.00
const rateEUR2RUB = rateUSD2RUB / rateUSD2EUR

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var amountEUR float64 = 100
	rateEUR2RUB := amountEUR * rateEUR2RUB
	println(rateEUR2RUB)
}
