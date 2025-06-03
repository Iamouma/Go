package main

import (
	"fmt"
)

func main() {
	i := true
	s := fmt.Sprintf("%v", i)
	fmt.Printf("Type : %T \nValue : %v\n\n", s, s)

	i1 := false
	s1 := fmt.Sprintf("%v", i1)
	fmt.Printf("Type : %T \nValue : %v\n", s1, s1)
}
