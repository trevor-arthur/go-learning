package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	fmt.Scan(&input)
	return input
}

func contains(menu []string, order string) bool {
	for _, item := range menu {
		if order == item {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)
	fmt.Println("Enter 'quit' to exit out of the menu.")

	var total int
	var order string
	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 4
		} else if order == "quit" {
			continue
		} else {
			fmt.Println("This item is not on the menu.")
		}
	}

	fmt.Printf("The total for your order is: $%d\n", total)

	for amount := total; amount > 0; amount -= 2 {
		if amount == total {
			fmt.Println("Uh-oh! Looks like I only have $2 bills.")
			fmt.Println("You hand the cashier a $2 bill.")
			continue
		} else {
			fmt.Printf("The remaining amount is: $%d\n", amount)
		}
		fmt.Println("You hand the cashier a $2 bill.")
	}
	fmt.Println("Thank you, have a great day!")
}
