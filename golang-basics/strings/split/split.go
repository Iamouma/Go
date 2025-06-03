package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Welcome,to,MotherEarth"
	fmt.Println("", s)

	result := strings.Split(s, ",")
	fmt.Println("Result:", result)
}
