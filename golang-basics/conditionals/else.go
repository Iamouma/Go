package main

import "fmt"

func main() {
	// if the condition in a if statement is not met then we execute the else
	// block if it is defined.
	isRaining := true

	if isRaining {
		fmt.Println("It's raining, take an umbrella")
	} else {
		fmt.Println("It's not raining. Umbrella not needed")
	}
}
