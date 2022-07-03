package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards! Good job, but remember, this the first step.")
	} else {
		isHeistOn = false
		fmt.Println("BUSTED: The guards caught you! Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("BUSTED: The vault can't be opened!")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Why did we rob the bank next to the police station?")
			fmt.Println("BUSTED: The cops were waiting for you outside the doors!")
		case 1:
			isHeistOn = false
			fmt.Println("Did somebody hear a door shut?")
			fmt.Println("BUSTED: The vault door closed behind you, locking you inside!")
		case 2:
			isHeistOn = false
			fmt.Println("Why won't the car start?!")
			fmt.Println("BUSTED: You made it to the getaway car, but it looks like someone forgot to fuel it up!")
		default:
			fmt.Println("Start the getaway car!")
			fmt.Println("SUCCESS: You made it to the getaway care and safely escaped!")
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Printf("Total haul from the heist: $%d\n", amtStolen)
		}
	}
}
