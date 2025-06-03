// Using the & operator to get the address of an existing struct.
package main

import "fmt"

type Person struct {
	name	string
	age	int
}

func main() {
	// creating an instance of the struct
	p1 := Person{name: "A", age: 25}

	// creating a pointer to the struct
	personPointer := &p1

	// accessing fields using the pointer
	fmt.Println("Name:", personPointer.name)
	fmt.Println("Age:", personPointer.age)

	// modifying struct values using the pointer
	personPointer.age = 26
	fmt.Println("Updated struct:", p1)
}


