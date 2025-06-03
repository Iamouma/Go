package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("", slice)
}

func main() {
	// initial slice
	a := []int{1, 2, 3, 4, 5}

	printSlice(a)
}
