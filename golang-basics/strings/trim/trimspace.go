package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " Hello, World   "
	result := strings.TrimSpace(s)
	fmt.Println(result)
}
