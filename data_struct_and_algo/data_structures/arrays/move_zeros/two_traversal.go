package main

import (
	"fmt"
)
func moveZerosToEnd(arr []int) {
	n := len(arr)
	count := 0 // index to place the next non-zero element

	// first traversal: move non-zero elements to the front
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	// second traversal: fill the rest with 0s
	for i := count; i < n; i++ {
		arr[i] = 0
	}
}
func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	moveZerosToEnd(arr)
	fmt.Println("Array after moving zeros to end:", arr) // Output: [1 2 4 3 5 0 0 0]
}
