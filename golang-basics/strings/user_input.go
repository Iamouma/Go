package main

import (
	"fmt"
	"strings"
)

func main() {
	// user input
	var userInput string
	fmt.Print("Enter string: ")
	fmt.Scanln(&userInput)

	// converting to lower case
	lowerInput := strings.ToLower(userInput)
	fmt.Println("Lower case input:", lowerInput)
}
