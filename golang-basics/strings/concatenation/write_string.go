package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 :="Hello, "
	s2 := "World!"

	// initializing a bytes buffer
	var b bytes.Buffer
	b.WriteString(s1)
	b.WriteString(s2)
	fmt.Println("", b.String())

}
