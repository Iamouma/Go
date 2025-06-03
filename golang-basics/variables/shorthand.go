package main

import "fmt"

func main() {
	// The := notation serves both as declaration and as initialization.
	// foo := "bar" is equivalent to var foo string = "bar"
	a := 9
	b := "golang"
	c := 4.17
	d := false
	e := "Hello"
	f := `Do you like golang, so far?`
	g := 'M'
	h := true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
