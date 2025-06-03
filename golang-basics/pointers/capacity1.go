// Go program to illustrate how to find the
// capacity of the pointer to an array
package main

import (
    "fmt"
)

// Main function
func main() {

    // Creating an array
    arr := [8]int{200, 300, 400,
       500, 600, 700, 100, 200}
    
    var x int

    // Creating pointer
    var p [5]*int

    // Assigning the address
    for x = 0; x < len(p); x++ {
        p[x] = &arr[x]
    }

    // Displaying result
    for x = 0; x < len(p); x++ {
    
        fmt.Printf("Value of p[%d] = %d\n",
                                 x, *p[x])
    }

    // Finding capacity
    // using cap() function
    fmt.Println("Capacity of arr: ", cap(arr))
    fmt.Println("Capacity of p: ", cap(p))
}
