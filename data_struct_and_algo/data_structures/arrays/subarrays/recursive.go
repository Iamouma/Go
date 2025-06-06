package main

import "fmt"

// Helper function to print a subarray from start to end index
func getSubarray(arr []int, start, end int) []int {
	sub := []int{}
	for i := start; i <= end; i++ {
		sub = append(sub, arr[i])
	}
	return sub
}

// Recursive function to generate subarrays
func generateSubarraysRecursive(arr []int, start, end int, result *[][]int) {
	n := len(arr)

	// Base condition: If end reaches end of array, stop recursion
	if end == n {
		return
	}

	// If start > end, reset start to 0 and increment end
	if start > end {
		generateSubarraysRecursive(arr, 0, end+1, result)
		return
	}

	// Append current subarray from start to end
	*result = append(*result, getSubarray(arr, start, end))

	// Recursive call with incremented start
	generateSubarraysRecursive(arr, start+1, end, result)
}

func main() {
	arr := []int{1, 2, 3}
	result := [][]int{}

	generateSubarraysRecursive(arr, 0, 0, &result)

	fmt.Println("All subarrays:")
	for _, sub := range result {
		fmt.Println(sub)
	}
}
