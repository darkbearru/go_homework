package main

const rateUSD2EUR = 1.14
const rateUSD2RUB = 69.00

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var amountEUR float64 = 100
	rateEUR2USD := 1 / rateUSD2EUR
	rateEUR2RUB := amountEUR * rateEUR2USD * rateUSD2RUB
	println(rateEUR2RUB)
}
