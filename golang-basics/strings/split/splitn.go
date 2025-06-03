package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Welcome,to,MotherEarth"
	fmt.Println("", s)

	result := strings.SplitN(s, ",", 2)
	fmt.Println("Result:", result)
}
