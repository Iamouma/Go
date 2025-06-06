package main

import (
	"fmt"
)

// findSecondLargest finds the second largest distinct element in one pass.
// Returns -1 if second largest does not exist.
func findSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest := -1
	secondLargest := -1

	for _, val := range arr {
		if val > largest {
			// Update secondLargest before updating largest
			secondLargest = largest
			largest = val
		} else if val < largest && val > secondLargest {
			// Update secondLargest only if it's not equal to the largest
			secondLargest = val
		}
	}

	return secondLargest
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

