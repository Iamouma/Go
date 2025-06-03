package main

import "fmt"

// Global variabale declaration
var myVariable int = 100

func main() {
	// Local variable inside the main function
	var localVar int = 200
	fmt.Printf("Inside main - Global variable: %d\n", myVariable)
	fmt.Printf("Inside main - Local variable: %d\n", localVar)
	display()
}
func display() {
	fmt.Printf("Inside display - Global variable: %d\n", myVariable)
}

