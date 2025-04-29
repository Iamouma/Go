package main

import "fmt"

func findMissingNumber(arr []int) int {
	n := len(arr) + 1
	total := (n * (n + 1)) / 2
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return total - sum
}

func main() {
	fmt.Println(findMissingNumber([]int{1, 2, 4, 5}))
}
