package main

import (
	"fmt"
	"strings"
)
func main() {
	s1 := "Hello, "
	s2 := "World!"

	// concatenating using strings.Join
	result := strings.Join([]string{s1, s2}, "")
	fmt.Println("", result)
}
