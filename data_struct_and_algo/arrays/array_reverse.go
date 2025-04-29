
package main

import "fmt"

func reverseArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(reverseArray([]int{1, 2, 3, 4})) // [4 3 2 1]
}
