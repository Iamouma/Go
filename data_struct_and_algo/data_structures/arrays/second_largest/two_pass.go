package main

import (
	"fmt"
	"math"
)

// findSecondLargest finds the second largest distinct element using two-pass search.
// Returns -1 if it doesn't exist.
func findSecondLargest(arr []int) int {
	n := len(arr)
	if n < 2 {
		return -1
	}

	// First Pass: Find the largest element
	max := math.MinInt64
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	// Second Pass: Find the largest element less than 'max'
	secondMax := math.MinInt64
	for _, val := range arr {
		if val != max && val > secondMax {
			secondMax = val
		}
	}

	// If no second max found, return -1
	if secondMax == math.MinInt64 {
		return -1
	}
	return secondMax
}

func main() {
	// Example usage
	arr := []int{12, 35, 1, 10, 34, 1}

	result := findSecondLargest(arr)
	if result != -1 {
		fmt.Println("Second Largest Element:", result)
	} else {
		fmt.Println("Second Largest Element does not exist")
	}
}

