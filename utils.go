package main

import (
	"fmt"
	"strconv"
)

// GetBet asks the user for a valid bet amount
func GetBet(balance int) int {
	for {
		fmt.Printf("Enter bet amount (Current balance: $%d) or 0 to quit: ", balance)
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		bet, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if bet == 0 {
			return 0
		}

		if bet < 0 {
			fmt.Println("Bet must be a positive number.")
			continue
		}

		if bet > balance {
			fmt.Println("You don't have enough balance for that bet!")
			continue
		}

		return bet
	}
}
