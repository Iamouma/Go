package main

import "fmt"

func main() {
	// Golang program to demonstrate the variables storing the hexadeciaml values
	// storing the hexadecimal values in variables.
	x := 0xFF
	y := 0x9C

	// Dispalying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)
}
