package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World!!"
	result := strings.Trim(s, "!!")
	fmt.Println(result)
}
