package main

import "fmt"

func main() {
	// creating a slice using the var keyword
	var my_slice_1 = []string{"Mario", "Yoshi", "Luigi"}
	fmt.Println("My Slice 1:", my_slice_1)

	// creating a slice using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
}
