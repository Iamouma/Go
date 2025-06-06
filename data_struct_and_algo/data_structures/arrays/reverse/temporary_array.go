package main

import "fmt"

// function to reverse the array using a temporary array
func reverseArrayNaive(arr []int) {
	n := len(arr)
	temp := make([]int, n)

	// copy elements in reverse order into temp
	for i := 0; i < n; i++ {
		temp[i] = arr[n-1-i]
	}

	// copy back to original array
	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}
}
func main() {
	arr := []int{1, 4, 3, 2, 6, 5}
	fmt.Println("Original array:", arr)

	reverseArrayNaive(arr)
	fmt.Println("Reversed array:", arr)
}
