// Golang program to illustrate bits.Sub() Function
package main

import (
    "fmt"
    "math/bits"
)

// Main function
func main() {

    // Finding diff and borrowOu
    // of the specified numbers
    // Using Sub() function
    nvalue_1, borrowOut := bits.Sub(4, 3, 0)
    fmt.Println("Diff:", nvalue_1)
    fmt.Println("BorrowOut :", borrowOut)
}
