// Golang program to demonstrate the use of type inference in pointer variables
package main

import "fmt"

func main() {
	// using var keyword, we are not defining any type with variable
	var y = 458

	// taking a pointer variable using var keyword without  specifying the type
	var p = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address f y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}
