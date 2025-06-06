package main

import "fmt"

// Recursive function to reverse array elements
func reverseRecursive(arr []int, left int, right int) {
	if left >= right {
		return
	}
	
	// Swap elements at positions left and right
	arr[left], arr[right] = arr[right], arr[left]
	
	// Recurse for the remaining subarray
	reverseRecursive(arr, left+1, right-1)
}

func main() {
	arr := []int{4, 5, 1, 2}
	fmt.Println("Original array:", arr)

	reverseRecursive(arr, 0, len(arr)-1)
	fmt.Println("Reversed array:", arr)
}

