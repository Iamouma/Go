// Go program to illustrate how to
// count the elements of the string
package main

import (
    "fmt"
    "strings"
)

// Main function
func main() {

    // Creating and initializing the strings
    str1 := "Welcome to Mother Earth"
    str2 := "My dog name is Dollar"
    str3 := "I like to play Ludo"

    // Displaying strings
    fmt.Println("String 1: ", str1)
    fmt.Println("String 2: ", str2)
    fmt.Println("String 3: ", str3)

    // Counting the elements of the strings
    // Using Count() function
    res1 := strings.Count(str1, "o")
    res2 := strings.Count(str2, "do")

    // Here, it also counts white spaces
    res3 := strings.Count(str3, "")

    // Displaying the result
    fmt.Println("\nResult 1: ", res1)
    fmt.Println("Result 2: ", res2)
    fmt.Println("Result 3: ", res3)
}
