package main

import "fmt"

func main() {
	// Strings In Go are defines with double quotes
	var s string = "This is string s"
	fmt.Println(s)

	// single quotes are not used to enable the useof apostrophes in strings without havig to escape
	var str string = "Don't worry about apostrophes"
	fmt.Println(str)
}
