// Go program to illustrate the 
// concept of comparing two pointers
package main

import "fmt"

func main() {

    val1 := 2345
    val2 := 567

    // Creating and initializing pointers
    var p1 *int
    p1 = &val1
    p2 := &val2
    p3 := &val1

    // Comparing pointers
    // with each other
    // Using != operator
    res1 := &p1 != &p2
    fmt.Println("Is p1 pointer not equal to p2 pointer: ", res1)

    res2 := p1 != p2
    fmt.Println("Is p1 pointer not equal to p2 pointer: ", res2)

    res3 := p1 != p3
    fmt.Println("Is p1 pointer not equal to p3 pointer: ", res3)

    res4 := p2 != p3
    fmt.Println("Is p2 pointer not equal to p3 pointer: ", res4)

    res5 := &p3 != &p1
    fmt.Println("Is p3 pointer not equal to p1 pointer: ", res5)
}
