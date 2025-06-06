package main

import (
	"fmt"
)

func findLeaders(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	leaders := []int{}
	maxFromRight := arr[n-1]
	leaders = append(leaders, maxFromRight) // last element is always a leader

	// Traverse from second last to first
	for i := n - 2; i >= 0; i-- {
		if arr[i] >= maxFromRight {
			maxFromRight = arr[i]
			leaders = append(leaders, arr[i])
		}
	}

	// Reverse to maintain order of appearance
	for i, j := 0, len(leaders)-1; i < j; i, j = i+1, j-1 {
		leaders[i], leaders[j] = leaders[j], leaders[i]
	}

	return leaders
}

func main() {
	fmt.Println(findLeaders([]int{16, 17, 4, 3, 5, 2})) // Output: [17 5 2]
	fmt.Println(findLeaders([]int{1, 2, 3, 4, 5, 2}))   // Output: [5 2]
}

