package main

import "fmt"

// Attempts to modify the value of num
func modify(num int) {
	num = 50
}

func main() {
	num := 20
	fmt.Printf("Before, num = %d\n", num)
	modify(num)
	fmt.Printf("After, num = %d\n", num)
}
