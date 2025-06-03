package main

import "fmt"

func main() {
	// Go has only one looping construct, the for loop
	// the basic for loop has three components separaeted by semicolons:
	// i.e * the init statement excuted before the first iteration.
	// * the condition expression evaluated before every iteration.
	// * the post statement executed at the end of every iteration.

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
