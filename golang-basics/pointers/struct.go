// We define a Person Struct with fields name and age. The variable p1 is an instance of this struct.
package main

import "fmt"

// defining a struct
type Person struct {
	name	string
	age	int
}

func main() {
	// creating an instance of the struct
	p1 := Person{name: "A", age: 25}

	fmt.Println("Original struct:", p1)
}
