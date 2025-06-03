package main

import "fmt"

// Variadic functions in Go allows you to pass a variable number of arguments to a function.
// This feature is useful when you don't know beforehand how many arguments you will pass.
// Variadic functions accept multiple arguments of the same type and can be called with any number of arguments.

// Variadic function to calculate sum.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 4, 5:", sum(4, 5))
	fmt.Println("Sum of no numbers:", sum())
}
