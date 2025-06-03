package main

import "fmt"

func modifySlice(slice []int) {

	for i := range slice {
		slice[i] *= 2// doubling each element
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", a)

	modifySlice(a)

	fmt.Println("Modified slice:", a)

}
