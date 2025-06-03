package main

import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func main() {
	// creating a destination array with the same size as the source array
	var destination [5]int

	// manually copying each element
	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}

	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}
