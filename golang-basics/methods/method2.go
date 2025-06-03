package main

import "fmt"

// Defining a struct
type person struct {
    name string
}

// Method with pointer receiver to modify data
func (p *person) changeName(newName string) {
    p.name = newName
}

func main() {
    a := person{name: "a"}
    
    fmt.Println("Before:", a.name)
    
    // Calling the method to change the name
    a.changeName("b")
    
    fmt.Println("After:", a.name)
}
