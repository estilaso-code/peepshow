package main

import (
	"fmt"

	"github.com/estilaso-code/goinput"
)

func greetUser() {
	fmt.Print(`
 _____                _____ _                   
|  __ \              / ____| |                  
| |__) |__  ___ _ __| (___ | |__   _____      __
|  ___/ _ \/ _ \ '_ \\___ \| '_ \ / _ \ \ /\ / /
| |  |  __/  __/ |_) |___) | | | | (_) \ V  V / 
|_|   \___|\___| .__/_____/|_| |_|\___/ \_/\_/  
               | |                              
               |_|                              
`, "Welcome to the show! Lets peep some passwords.\n")
}

var complexity string

func passStrength() error {
	input, err := goinput.Input("How hard do you want this password?\nType 1 for: So good and hard\nType 2 for: Kinda mid\nType 3 for: FLacid like your dick!\n")
	if err != nil {
		return fmt.Errorf("User can't follow instructions: %w", err)
	}
	complexity = input
	return nil
}

func complexityChoice(complexity string) {
	switch complexity {
	case "1":
		fmt.Println("Oh, you a big boy huh?")
	case "2":
		fmt.Println("Meh, kinda pansy, but whatever bro")
	case "3":
		fmt.Println("You little sissy boy! Grow some balls!!")
	default:
		fmt.Println("You are a real dumbass.")
	}

}

func main() {
	greetUser()
	passStrength()
	complexityChoice(complexity)
}
