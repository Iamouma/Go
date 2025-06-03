package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// creating and initializing a string using shorthand declaration
	mystr := "Welcome to Mother Earth"

	// finding the length of the string using len() function
	length1 := len(mystr)

	// using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystr)

	// displaying the length of the string
	fmt.Println("string:", mystr)
	fmt.Println("Length 1:", length1)
	fmt.Println("Length 2:", length2)

}
