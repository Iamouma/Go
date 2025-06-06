package main

import "fmt"

func isSortedRecursive(arr []int, index int) bool {
	// base case: 0 or 1 element is sorted
	if index == len(arr)-1 {
		return true
	}

	// if current is greater than next, not sorted
	if arr[index] > arr[index+1] {
		return false
	}
	// check next
	return isSortedRecursive(arr, index+1)
}
func main() {
	fmt.Println(isSortedIterative([]int{20, 21, 45, 89, 89, 90})) // Yes
	fmt.Println(isSortedIterative([]int{20, 20, 45, 89, 89, 90})) // Yes
	fmt.Println(isSortedIterative([]int{20, 20, 78, 98, 99, 97})) // No
}
