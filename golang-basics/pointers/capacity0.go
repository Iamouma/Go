// Go program to illustrate how to find the
// capacity of the pointer to an array
package main

import (
    "fmt"
)

// Main function
func main() {

    // Creating and initializing
    // pointer to array
    // Using var keyword
    var ptr1 [7]*int
    var ptr2 [5]*string
    var ptr3 [8]*float64

    // Finding the capacity of
    // the pointer to array
    // Using cap function
    fmt.Println("Capacity of ptr1: ", cap(ptr1))
    fmt.Println("Capacity of ptr2: ", cap(ptr2))
    fmt.Println("Capacity of ptr3: ", cap(ptr3))

}
