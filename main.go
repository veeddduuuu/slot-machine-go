package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getName() string {
	name := ""
	fmt.Println("=============================")
	fmt.Println("    WELCOME TO THE CASINO    ")
	fmt.Println("=============================")
	fmt.Printf("Please enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	fmt.Printf("\nHello %s, welcome to the Slot Machine!\n", name)
	return name
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	name := getName()
	if name == "" {
		return
	}

	balance := 100 // Starting balance

	for {
		if balance <= 0 {
			fmt.Println("\nYou are out of money! Game over.")
			break
		}

		fmt.Println("-----------------------------")
		bet := GetBet(balance)
		if bet == 0 {
			fmt.Printf("\nThanks for playing, %s! You are cashing out with $%d.\n", name, balance)
			break
		}

		// Deduct bet from balance
		balance -= bet

		// Spin the slot machine
		fmt.Println("\nSpinning...")
		time.Sleep(1 * time.Second) // add a little suspense

		spinResult := Spin()
		fmt.Printf("\n[ %s | %s | %s ]\n\n", spinResult[0], spinResult[1], spinResult[2])

		// Calculate payout
		winnings := CalculatePayout(spinResult, bet)
		if winnings > 0 {
			fmt.Printf("WINNER! You won $%d!\n", winnings)
			balance += winnings
		} else {
			fmt.Println("No luck this time. You lost your bet.")
		}
	}
}