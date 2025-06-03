package main

import "fmt"

func main() {
	s1 := "Hello, "
	s2 := "World!"

	// concatenating using fmt.Sprintf
	result := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println("",result)
}
