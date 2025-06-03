// A simple Go program to Illustrate how to create a rune
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// creating a rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2, rune2, reflect.TypeOf(rune2))
	fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3, reflect.TypeOf(rune3))

}
