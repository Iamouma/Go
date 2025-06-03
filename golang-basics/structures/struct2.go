// Pointers in Go programming language is a variable which is used to store the memoty address of anther varable.

package main

import "fmt"

// defining a struct
type Employee struct {
	firstName, LastName string
	age, salary int
}

func main() {
	// passing the address of struct variable
	// emp8 isa to the Employee struct
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

	// (*emp8).firstName is the syntax to access
    	// the firstName field of the emp8 struct
	fmt.Println("First Name:", (*emp8).firstName)
    	fmt.Println("Age:", (*emp8).age)
}
