// Go program to illustrate how to check
// the string is present or not in the
// specified string
package main

import (
    "fmt"
    "strings"
)

// Main function
func main() {

    // Creating and initializing strings
    str1 := "Welcome to Mother Earth"
    str2 := "Finding and Interacting with nature"

    fmt.Println("Original strings")
    fmt.Println("String 1: ", str1)
    fmt.Println("String 2: ", str2)

    // Checking the string present or not
    //  Using Contains() function
    res1 := strings.Contains(str1, "Earth")
    res2 := strings.Contains(str2, "World")

    // Displaying the result
    fmt.Println("\nResult 1: ", res1)
    fmt.Println("Result 2: ", res2)

}
