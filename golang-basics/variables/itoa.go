// Go program to illustrate how to convert integer variable into String
package main

import (
	"fmt"
	"strconv"
)
func main() {
	i := 32
	s := strconv.Itoa(i)
	fmt.Printf("Type : %T \nValue : %v", s, s)
}
