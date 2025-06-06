package main

import "fmt"

func findMaxRecursive(arr []int, n int) int {
	if n == 1 {
		return arr[0]
	}

	max := findMaxRecursive(arr, n-1)

	if arr[n-1] > max {
		return arr[n-1]
	}
	return max
}
func main() {
	arr := []int{10, 20, 4}

	fmt.Println("Recursive Max:", findMaxRecursive(arr, len(arr)))
}
