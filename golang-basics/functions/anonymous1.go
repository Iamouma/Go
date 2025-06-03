// You can assign an anonymous function to a variable.
// This variable can be called like a regular function

package main

import "fmt"

func main() {
	// Assigning an anonymous function to a variable
	value := func() {
		fmt.Println("Hello World!")
	}

	value()
}
