package main

import (
	"fmt"
)

func main() {
	var currencies = map[string]float32{
		"JPY": 136.14,
		"EUR": 0.96,
		"CAD": 1.29,
		"RUB": 55.03,
		"GPR": 0.83,
		"CNY": 6.69,
		"MXN": 20.28,
	}
	currencyNames := []string{}
	for currencyName := range currencies {
		currencyNames = append(currencyNames, currencyName)
	}
	fmt.Println("Welcome to Currency Coverter!")
	var dollarAmount float32
	var currency string
	fmt.Print("Please enter dollar amount: $")
	fmt.Scan(&dollarAmount)
	if dollarAmount == 0 {
		fmt.Println("Error: Unknown dollar amount.")
	} else {
		fmt.Println("Available convertable currencies:", currencyNames)
		fmt.Print("Please enter desired currency: ")
		fmt.Scan(&currency)
		rate, isValid := currencies[currency]
		if isValid == true {
			convertedAmount := rate * dollarAmount
			fmt.Println("The converted amount is:", convertedAmount)
		} else {
			fmt.Println("Error: Unknown currency type.")
		}
	}
}
