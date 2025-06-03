// Golang program to demonstrate the
// use of Named Return Arguments

package main

import "fmt"

// Main Method
func main() {

    // calling the function, here
    // function returns two values
    m, d := calculator(105, 7)

    fmt.Println("105 x 7 = ", m)
    fmt.Println("105 / 7 = ", d)
}

// function having named arguments
func calculator(a, b int) (mul int, div int) {

    // here, simple assignment will
    // initialize the values to it
    mul = a * b
    div = a / b

    // here you have return keyword
    // without any resultant parameters
    return
}
