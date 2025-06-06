package main

import (
	"fmt"
	"sort"
)

// findSecondLargest finds the second largest distinct element in ana array
// Returns -1 if the second largest doesn't exist

func findSecondLargest(arr []int) int {
	n := len(arr)
	if n < 2 {
		return -1
	}

	// sort the array in ascending order
	sort.Ints(arr)

	// Start from the second last element and move backward
	largest := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] < largest {
			return arr[i]
		}
	}

}

func main() {
	// Example input
	arr := []int{12, 35, 1, 10, 34, 1}

	result := findSecondLargest(arr)
	if result != -1 {
		fmt.Println("Second Largest Element:", result)
	} else {
		fmt.Println("Second Largest Element does not exist")
	}
}
