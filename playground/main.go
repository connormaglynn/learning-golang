package main

import "fmt"

func main() {
	// Add program to execute
	guessTheNumber(2)
}

func guessTheNumber(guess int) bool {
	answer := 1
	if guess == answer {
		fmt.Print(guess, " is correct! 🥳")
		return true
	}
	fmt.Print(guess, " is wrong 😭 try again! 1️⃣🤫👀")
	return false
}
