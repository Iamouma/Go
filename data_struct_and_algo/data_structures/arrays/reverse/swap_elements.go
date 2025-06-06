package main

import "fmt"

func reverseArray(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		// swap
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}
func main() {
	arr := []int{4, 5, 1, 2}
	fmt.Println("Original array:", arr)

	reverseArray(arr)
	fmt.Println("Reversed array:", arr)
}
