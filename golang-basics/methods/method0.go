// Go methods are like functions but with a key difference; they have a receiver argument, which alows access to the receivers properties.
// The receiver can be a struct or non-struct type , but both must be inthe same package.
// Methods cannot be create for types defines in other packages, including built-in types like int or string.
package main

import "fmt"

// Defining a struct
type person struct {
	name string
	age int
}

// Defining a method with struct reeiver
func (p person) display() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func main() {
	// creating an instance of the struct
	a := person{name: "a", age: 25}

	// calculate the method
	a.display()
}
