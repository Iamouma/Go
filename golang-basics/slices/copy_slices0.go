package main

import "fmt"

var source = []int{10, 20, 30, 40, 50}

func main() {
	// creating a destination slice with the same length as the source
	destination := make([]int, len(source))

	// copying the elements from source to destination
	count := copy(destination, source)

	// printing the slices
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
	fmt.Println("Elementscopied:", count)
}
