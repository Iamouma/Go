package main

import "fmt"

// Modifies value of num via reference
func modify(num *int) {
	*num = 50
}

func main() {
	num := 20
	fmt.Printf("Before, num = %d\n", num)
	modify(&num)
	fmt.Printf("After, num = %d\n", num)
}
