package main

import "fmt"

func main() {
	// creating and initializing a string
	str := "Welcome to Mother Earth"

	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {
		fmt.Println("\nCharacter = %c Bytes = %x", str[c], str[c])
	}
}
