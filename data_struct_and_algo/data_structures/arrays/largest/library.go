package main

import (
	"fmt"
	"sort"
)
func findMaxUsingLibrary(arr []int) int {
	if len(arr) == 0 {
		panic("Array is empty")
	}
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	return sortedArr[len(sortedArr)-1]
}
func main() {
	arr := []int{10, 20, 4}

	fmt.Println("Library Max:", findMaxUsingLibrary(arr))
}
