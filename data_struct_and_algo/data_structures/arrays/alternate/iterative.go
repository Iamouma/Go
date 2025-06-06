package main

import "fmt"

// Itrative Approach
// Print alternate elements starting from index 0, incrementing by 2
func getAlternates(arr []int) {
	// Iterate over all alternate elements
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	// Test case 1
	// Expected output: 10, 20, 30
	arr1 := []int{10, 20, 30, 40, 50}
	fmt.Print("Input: ", arr1, " Output: ")
	getAlternates(arr1)

	// Tast case 2
	// Expected output: -5, 4, 12
	arr2 := []int{-5, 1, 4, 2, 12}
	fmt.Print("Input: ", arr2, " Output: ")
	getAlternates(arr2)
} 
