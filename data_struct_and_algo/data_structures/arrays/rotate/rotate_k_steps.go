package main

import "fmt"

func rotateArray(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	return append(arr[n-k:], arr[:n-k]...)
}

func main() {
	fmt.Println(rotateArray([]int{11, 12, 13, 14, 15}, 2))
}
