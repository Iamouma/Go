package main

import "fmt"

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}))
}
