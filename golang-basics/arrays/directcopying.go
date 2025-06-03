package main

import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func main() {
	// copying by direct assignment
	var destination [5]int = source

	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}
