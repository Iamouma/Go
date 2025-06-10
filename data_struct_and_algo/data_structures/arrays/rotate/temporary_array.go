package main

import (
	"fmt"
)

func rotateRight(arr []int, d int) []int {
	n := len(arr)
	d = d % n // In case d > n

	temp := make([]int, n)

	// Step 1: Copy last d elements to the beginning of temp
	for i := 0; i < d; i++ {
		temp[i] = arr[n-d+i]
	}

	// Step 2: Copy first (n - d) elements to the end of temp
	for i := d; i < n; i++ {
		temp[i] = arr[i-d]
	}

	// Step 3: Copy temp back to original array
	copy(arr, temp)
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	d := 2
	result := rotateRight(arr, d)
	fmt.Println("Rotated array:", result) // Output: [4 5 1 2 3]
}

