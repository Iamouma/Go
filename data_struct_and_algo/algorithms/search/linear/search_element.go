package main

import (
	"fmt"
)

// linearSearch performs a linear search on the way.
// It returns the index of the first occurrence of x or -1 if not 
func linearSearch(arr []int, x int) int {
	for i, value := range arr {
		if value == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	x := 30

	result := linearSearch(arr, x)

	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", x, result)
	} else {
		fmt.Printf("Element %d not found int the array\n", x)
	}
}
