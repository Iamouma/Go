// Go program to illustrate how to convert Boolean Type into String
package main

import (
	"fmt"
	"strconv"
)

// main function
func main() {
	i := true
	s := strconv.FormatBool(i)
	fmt.Printf("Type : %T \nValue : %\n", s, s)
}
