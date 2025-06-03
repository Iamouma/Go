package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "@@Hello, World"
	result := strings.TrimSuffix(s, "orld")
	fmt.Println(result)
}
