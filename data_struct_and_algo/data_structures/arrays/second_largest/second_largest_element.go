package main

import (
	"fmt"
	"math"
)

func secondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest := math.MinInt64
	second := math.MinInt64

	for _, value := range arr {
		if value > largest {
			second = largest
			largest = value
		} else if value > second && value != largest {
			second = value
		}
	}
	
	if second == math.MinInt64 {
		return -1
	}
	return second
}

func main() {
	fmt.Println(secondLargest([]int{10, 20, 20, 4, 45, 45})) // 20
	fmt.Println(secondLargest([]int{5, 5, 5}))               // -1
	fmt.Println(secondLargest([]int{10}))                    // -1
}
