package main

import (
	"fmt"
)

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotateRight(arr []int, d int) []int {
	n := len(arr)
	d = d % n // In case d > n

	// Step 1: Reverse the entire array
	reverse(arr, 0, n-1)

	// Step 2: Reverse first d elements
	reverse(arr, 0, d-1)

	// Step 3: Reverse remaining n-d elements
	reverse(arr, d, n-1)

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	d := 2
	result := rotateRight(arr, d)
	fmt.Println("Rotated array:", result) // Output: [4 5 1 2 3]
}

