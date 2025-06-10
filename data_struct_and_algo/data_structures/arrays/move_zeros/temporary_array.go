package main

import (
	"fmt"
)
func moveZerosToEnd(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	temp := make([]int, n)
	index := 0

	// copy non-zero elements to temp
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			temp[index] = arr[i]
			index++
		}
	}

	// fill the rest with zeros
	for index < n {
		temp[index] = 0
		index++
	}

	copy temp array back to original array
	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}
}

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	moveZerosToEnd(arr)
	fmt.Println("Array after moving zeros to end:", arr) // Output: [1 2 4 3 5 0 0 0]
}
