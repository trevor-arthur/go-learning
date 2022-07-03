package main

import (
	"fmt"
)

// A space travel simulator created to test knowledge of Go functions.
// https://github.com/trevor-arthur

func fuelGauge(fuel int) {
	fmt.Println("Fuel Gauge:", fuel, "gallons.")
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Print("We have arrived! Welcome to ", planet, "!\n")
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func navigationMenu() string {
	fmt.Println("Navigation Menu:")
	fmt.Println("- Venus")
	fmt.Println("- Mercury")
	fmt.Println("- Mars")
	fmt.Print("Where do you want to fly to? ")
	var destination string
	fmt.Scan(&destination)
	fmt.Printf("Target Destination: %v\n", destination)
	return destination
}

func main() {
	fmt.Println("Welcome to your Spaceship!")
	var fuel int = 1000000
	var planetChoice string
	fuelGauge(fuel)
	planetChoice = navigationMenu()
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
