package main

import "fmt"

func reverseArray(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		// swap elemts at left and right
		arr[left], arr[right] = arr[right], arr[left]

		// move pointers toward the center
		left++
		right--
	}
}
func main() {
	arr := []int{1, 4, 3, 2, 6, 5}
	fmt.Println("Original array:", arr)

	reverseArray(arr)
	fmt.Println("Reversed array:", arr)
}
