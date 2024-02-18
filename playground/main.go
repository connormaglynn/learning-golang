package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Add program to execute
	//guessTheNumber(2)
	pointersAndReferences()
	//screamBackAtMe()
	// arrays()
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

	fmt.Println("")

	a := "👑"
	fmt.Println("a := 👑")
	fmt.Println("a ->", a, `// 👈 The variable "a" contains the string "👑"`)
	fmt.Println("&a ->", &a, `// 👈 The address of the variable "a"`)
	fmt.Println(`*a -> ERROR // 👈 Cannot be done since the variable "a" is not a pointer type (pointer type -> &)`)

	fmt.Println("")

	b := &a
	fmt.Println("b := &a")
	fmt.Println("b ->", b, `// 👈 The value of the variable "b" is the address of variable "a"`)
	fmt.Println("&b ->", &b, `// 👈 The address of the variable "b"`)
	fmt.Println("*b ->", *b, `// 👈 Go to the address stored in variable "b" and return the value (de-referencing)`)

	fmt.Println("")

	c := &b
	fmt.Println("c := &b")
	fmt.Println("c ->", c, `// 👈 The value of the variable "c" is the address of variable "b"`)
	fmt.Println("&c ->", &c, `// 👈 The address of the variable "c"`)
	fmt.Println("*c ->", *c, `// 👈 Go to the address stored in variable "c" and return the value (de-referencing)`)
	fmt.Println("**c ->", **c, `// 👈 Go to the address, of the address stored in variable "c" and return the value (de-referencing twice!)`)
	fmt.Println("")
	fmt.Println("End")
}

func screamBackAtMe() {
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.ToUpper(strings.TrimSpace(s))
	fmt.Println(s + "!")
}

func arrays() {
	fmt.Println("\nInit empty array and populate a single value by index")
	names := [4]string{}
	names[0] = "Connor"
	fmt.Println(names)

	fmt.Println("\nInit empty array and populate a multiple value by index")
	colours := [4]string{}
	colours[0] = "red"
	colours[2] = "blue"
	fmt.Println(colours)

	fmt.Println("\nInit array with values, overriding one by index")
	cities := [4]string{"sheffield", "doncaster", "london", "leeds"}
	cities[3] = "replaced!"
	fmt.Println(cities)

	fmt.Println("\nCopy cities array by value, modify the copy and display both")
	citiesCopiedByValue := cities
	citiesCopiedByValue[0] = "replaced!"
	fmt.Println("Original: ", cities)
	fmt.Println("Copy: ", citiesCopiedByValue)

	fmt.Println("\nCopy cities array by reference, modify the copy and display both (no de-referencing on copy)")
	citiesCopiedByReference := &cities
	citiesCopiedByReference[0] = "replaced!"
	fmt.Println("Original: ", cities)
	fmt.Println("Copy: ", citiesCopiedByReference)

	fmt.Println("\nCopy colours array by reference, modify the copy and display both  (with de-referencing on copy)")
	coloursCopiedByReference := &colours
	coloursCopiedByReference[0] = "replaced!"
	fmt.Println("Original: ", colours)
	fmt.Println("Copy: ", *coloursCopiedByReference)
}
