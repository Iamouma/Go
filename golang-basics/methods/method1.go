package main

import "fmt"

// Creating a custom type based on int
type number int

// Defining a method with a non-struct receiver
func (n number) square() number {
    return n * n
}

func main() {
    a := number(4)
    b := a.square()

    fmt.Println("Square of", a, "is", b)
}
