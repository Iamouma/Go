package main

import "fmt"

func main() {
	// let's try again changing the username variable
	userName := "Mathew"

	// prints You are not John
	if userName == "John" {
	    fmt.Println("You are John")
	} else {
	    fmt.Println("You are not John")
	}
}
