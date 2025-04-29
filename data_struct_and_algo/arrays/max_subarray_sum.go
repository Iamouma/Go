package main

import "fmt"

func maxSubArraySum(arr []int) int {
	maxSoFar := arr[0]
	currentMax := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > currentMax+arr[i] {
			currentMax = arr[i]
		} else {
			currentMax += arr[i]
		}
		if currentMax > maxSoFar {
			maxSoFar = currentMax
		}
	}
	return maxSoFar
}

func main() {
	fmt.Println(maxSubArraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
