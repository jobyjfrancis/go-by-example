package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* Variables Declaration
	Declare all variables needed for the Heist */
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	fmt.Println(eludedGuards, "\n")

	/* Prompting Heist, checking if Heist is possible */

	if eludedGuards >= 50 {
		fmt.Println("isHeistOn:", isHeistOn, "\n")
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.\n")

	} else {
		isHeistOn := false
		fmt.Println("isHeistOn:", isHeistOn, "\n")
		fmt.Println("Plan a better disguise next time\n")
	}

	openedVault := rand.Intn(100)

	fmt.Println(openedVault, "\n")

	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and Go!\n")
	} else if isHeistOn == false {
		fmt.Println("The vault can't be opened\n")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn := false
			fmt.Println(isHeistOn, "Heist failed")
		case 1:
			isHeistOn := false
			fmt.Println(isHeistOn, "Turns out the vault door won't open from the inside")
		case 2:
			isHeistOn := false
			fmt.Println(isHeistOn, "Security in front of the exit door")
		case 3:
			isHeistOn := false
			fmt.Println(isHeistOn, "Sensors in the doorway")
		default:
			fmt.Println(isHeistOn, "Heist Success, Start the getaway car!")
		}

		if isHeistOn == true {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println("\nThe amount stolen is: ", amtStolen)
		}
	}

	fmt.Println("\nisHeistOn:", isHeistOn)

}
