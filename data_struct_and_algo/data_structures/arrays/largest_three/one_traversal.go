package main

import (
	"fmt"
	"math"
)

func topThreeDistinct(arr []int) []int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range arr {
		// skip if this number is already one of the top three
		if num == first || num == second || num == third {
			continue
		}

		if num > first {
			third = second
			second = first
			first = num
		} else if num > second {
			third = second
			second = num
		} else if num  > third {
			third = num
		}
	}
	
	// Collect results
	result := []int{}
	if first != math.MinInt64 {
		result = append(result, first)
	}
	if second != math.MinInt64 {
		result = append(result, second)
	}
	if third != math.MinInt64 {
		result = append(result, third)
	}

	return result
}
func main() {
	fmt.Println(topThreeDistinct([]int{10, 4, 3, 50, 23, 90})) // [90 50 23]
	fmt.Println(topThreeDistinct([]int{10, 9, 9}))             // [10 9]
	fmt.Println(topThreeDistinct([]int{10, 10, 10}))           // [10]
	fmt.Println(topThreeDistinct([]int{}))                     // []
}
