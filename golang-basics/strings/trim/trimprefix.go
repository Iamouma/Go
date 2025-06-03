package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "@@Hello, World!!"
	result := strings.TrimPrefix(s, "@@")
	fmt.Println(result)
}
