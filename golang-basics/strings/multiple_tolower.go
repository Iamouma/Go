package main

import (
	"fmt"
	"strings"
)

func main() {
	// initil strings
	phrases := []string{"Hello, World!", "Golang is Fun!"}

	fmt.Println("", phrases)

	// converting to lower case
	for i, s := range phrases {
		phrases[i] = strings.ToLower(s)
	}
	fmt.Println("", phrases)
}
