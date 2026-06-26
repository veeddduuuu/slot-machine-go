package main

import (
	"math/rand"
)

var symbols = []string{"🍒", "🍋", "🍉", "⭐", "🎰"}

// Spin generates a random 1-line spin with 3 symbols
func Spin() []string {
	result := make([]string, 3)
	for i := 0; i < 3; i++ {
		// Randomly select a symbol
		result[i] = symbols[rand.Intn(len(symbols))]
	}
	return result
}

// CalculatePayout checks if the spin was a winning spin and calculates payout
func CalculatePayout(spinResult []string, bet int) int {
	// Must match all 3
	if spinResult[0] == spinResult[1] && spinResult[1] == spinResult[2] {
		symbol := spinResult[0]
		
		multiplier := 0
		switch symbol {
		case "🍒":
			multiplier = 2
		case "🍋":
			multiplier = 3
		case "🍉":
			multiplier = 5
		case "⭐":
			multiplier = 10
		case "🎰":
			multiplier = 50
		}
		
		return bet * multiplier
	}
	
	// If no match, return 0
	return 0
}
