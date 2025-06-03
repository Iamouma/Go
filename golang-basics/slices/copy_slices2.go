package main

import "fmt"

var source = [] int{10, 20, 30, 40, 50}

func main() {
	// copying using a slice literal
	destination := []int{}
	destination = append(destination, source...)

	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
}
