// this function allocates memeory and returns a pointer to the struct
package main

import "fmt"

// defining a struct
type Person struct {
	name	string
	age	int
}

func main() {
	// creating a pointer to a new instance of Person
	personPointer := new(Person)
	personPointer.name = "B"
	personPointer.age = 30

	fmt.Println("Struct created with new:", *personPointer)
}
