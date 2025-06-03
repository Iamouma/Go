// A struct in Golang ia a user defined type that allows to group/combine items of possibly different types into a single type 
// Any real world entity which has some set of properties/fields can be represented as a struct.
package main

import "fmt"

// defining a struct
type Address struct {
	Name	string
	city	string
	Pincode	int
}

func main() {
	// Declaring a variable of a struct type
	// All the struct fields are initialized with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a struct using a struct literal
	a1 := Address{"Akshay", "Dehradun", 36235772}
	fmt.Println("Address1: ", a1)

	// Naming fields while initializing a struct
	a2 := Address{Name: "Anikaa", city: "Ballia", Pincode: 277001}
	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)
}
