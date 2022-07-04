package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomElement(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}

func main() {
	guests := []string{"Gary", "Erica", "Shay"}
	places := []string{"bedroom", "bathroom", "garage"}
	fmt.Println("A number of guests are attending a party in your parlor when the lights go out. You hear an angry “MEOW!” When the lights come back on, you realize — your cat is gone! It’s up to you to find out who hid your cat and where the cat is hidden!")

	fmt.Printf("It must have been %v, %v, or %v.\n", guests[0], guests[1], guests[2])
	var catnapperGuess string
	fmt.Print("Who do you think is the catnapper? ")
	fmt.Scan(&catnapperGuess)

	fmt.Printf("They could have hidden the cat in the %v, the %v, or the %v.\n", places[0], places[1], places[2])
	var catLocationGuess string
	fmt.Print("Where do you think the cat is? ")
	fmt.Scan(&catLocationGuess)

	rand.Seed(time.Now().UnixNano())
	catnapper := getRandomElement(guests)
	catLocation := getRandomElement(places)

	if catnapperGuess == catnapper && catLocationGuess == catLocation {
		fmt.Println("You were right! You caught the catnapper! You're a hero!")
	} else if catLocationGuess == catLocation {
		fmt.Printf("You found your cat in the %v, but you have no idea who put him there!\n", catLocation)
	} else {
		fmt.Printf("After a little while you hear someone burst out in laughter. It turns out %v hid the cat in the %v as a prank!\n", catnapper, catLocation)
	}
}
