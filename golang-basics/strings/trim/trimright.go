package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "@@Hello, World"
	result := strings.TrimRight(s, "World")
	fmt.Println(result)
}
