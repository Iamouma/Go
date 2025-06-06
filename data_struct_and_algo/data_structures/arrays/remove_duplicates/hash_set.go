package main

import (
	"fmt"
	"sort"
)

func removeDuplicatesWithSet(arr []int) []int {
	seen := make(map[int]bool)
	res := []int{}
	for _, val := range arr {
		if !seen[val] {
			res = append(res, val)
			seen[val] = true
		}
	}
	sort.Ints(res)
	return res
}
func main() {
	fmt.Println(removeDuplicatesWithSet([]int{4, 1, 2, 4, 3, 2, 1})) // [1 2 3 4]
}
