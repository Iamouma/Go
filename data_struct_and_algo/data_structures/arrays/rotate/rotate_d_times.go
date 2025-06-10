package main

import (
	"fmt"
)

func rotateRight(arr []int, d int) []int {
	n := len(arr)
	for i := 0; i < d; i++ {
		last := arr[n-1]
		for j := n - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = last
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	d := 2
	result := rotateRight(arr, d)
	fmt.Println("Rotated array:", result) // Output: [4 5 1 2 3]
}

