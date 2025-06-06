package main

import "fmt"

func generateSubarrays(arr []int) [][]int {
	n := len(arr)
	result := [][]int{}

	// Outermost loop - start index
	for start := 0; start < n; start++ {
		// Middle loop - end index
		for end := start; end < n; end++ {
			subarray := []int{}
			// Innermost loop - build the subarray
			for k := start; k <= end; k++ {
				subarray = append(subarray, arr[k])
			}
			result = append(result, subarray)
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3}
	fmt.Println("Subarrays of", arr1, "are:")
	for _, sub := range generateSubarrays(arr1) {
		fmt.Println(sub)
	}

	arr2 := []int{1, 2}
	fmt.Println("\nSubarrays of", arr2, "are:")
	for _, sub := range generateSubarrays(arr2) {
		fmt.Println(sub)
	}
}
