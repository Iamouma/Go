package main

import (
	"fmt"
)

func moveZerosToEnd(arr []int) {
	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			// Swap current element with element at count
			arr[i], arr[count] = arr[count], arr[i]
			count++
		}
	}
}

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	moveZerosToEnd(arr)
	fmt.Println("Array after moving zeros to end:", arr) // Output: [1 2 4 3 5 0 0 0]
}

