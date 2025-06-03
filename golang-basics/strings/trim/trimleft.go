package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "@@Hello, World!!"
	result := strings.TrimLeft(s, "@!")
	fmt.Println(result)
}
