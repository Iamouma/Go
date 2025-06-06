package main

import "fmt"

func removeDuplicatesSorted(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Pointer to last unique element
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i] // Move unique element forward
		}
	}
	return j + 1 // Length of unique part
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	k := removeDuplicatesSorted(arr)
	fmt.Println("Output:", arr[:k]) // [1 2 3 4 5]

	arr2 := []int{2, 2, 2, 2, 2}
	k2 := removeDuplicatesSorted(arr2)
	fmt.Println("Output:", arr2[:k2]) // [2]

	arr3 := []int{1, 2, 3}
	k3 := removeDuplicatesSorted(arr3)
	fmt.Println("Output:", arr3[:k3]) // [1 2 3]
}

