package main

import "fmt"

func isSortedIterative(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isSortedIterative([]int{20, 21, 45, 89, 89, 90})) // Yes
	fmt.Println(isSortedIterative([]int{20, 20, 45, 89, 89, 90})) // Yes
	fmt.Println(isSortedIterative([]int{20, 20, 78, 98, 99, 97})) // No
}
