// Golang program to illustrate bits.TrailingZeros64() Function
package main

import (
    "fmt"
    "math/bits"
)

// Main function
func main() {

    // Using TrailingZeros64() function
    a := bits.TrailingZeros64(15)
    fmt.Printf("Total number of trailing"+
            " zero bits in %d: %d", 15, a)
}
