package main

import (
	"fmt"
)

func findMaxIterative(arr []int) int {
	if len(arr) == 0 {
		panic("Array is empty")
	}

	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	arr := []int{10, 20, 4}

	fmt.Println("Iterative Max:", findMaxIterative(arr))
}
