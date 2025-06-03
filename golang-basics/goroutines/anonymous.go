package main

import (
	"fmt"
	"time"
)
func main() {
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello from Anonymous Goroutine!")

	time.Sleep(2 * time.Second) // Allow Goroutine to finish
	fmt.Println("Main function complete.")
}
