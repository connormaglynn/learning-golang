package main

import "fmt"

func main() {
	// Add program to execute
	//guessTheNumber(2)
	pointersAndReferences()
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

func pointersAndReferences() {
	fmt.Println("Start")

	a := "a"
	b := &a

	fmt.Println("a = ", a, `// 👈 The variable "a" contains the string "a"`)
	fmt.Println("&a = ", &a, `// 👈 The address of the variable "a"`)
	fmt.Println(`*a =  ERROR // 👈 Cannot be done since the variable "a" is not a pointer type (pointer type = &)`)
	fmt.Println("b = ", b, `// 👈 The value of the variable "b" is the address of variable "a"`)
	fmt.Println("&b = ", &b, `// 👈 The address of the variable "b"`)
	fmt.Println("*b = ", *b, `// 👈 Go to the address stored in variable "b" and return the value (de-referencing)`)

	fmt.Println("End")
}
