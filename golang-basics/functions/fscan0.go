// Golang program to illustrate the usage of
// fmt.Fscan() function

// Including the main package
package main

// Importing fmt, io and strings
import (
    "fmt"
    "os"
    "strings"
)

// Calling main
func main() {

        // Declaring some type of variables
    var (
        i int
        b bool
        s string
    )
    
    // Calling the NewReader() function to
    // specify some type of texts.
    // variable "r" contains the scanned texts
    r := strings.NewReader("10 false GFG")
    
    // Calling the Fscan() function to receive 
    // the scanned texts
    n, err := fmt.Fscan(r, &i, &b, &s)
    
    // If the above function returns an error then
    // below statement will be executed
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
    }
    
    // Printing each type of scanned texts
    fmt.Println(i, b, s)
    
    // It returns the number of items 
    // successfully scanned
    fmt.Println(n)
}
