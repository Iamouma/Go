// Golang program to illustrate the usage of
// fmt.Fscanln() function

// Including the main package
package main

// Importing fmt, io and strings
import (
    "fmt"
    "io"
    "strings"
)

// Calling main
func main() {

    // Declaring list of strings,
    // integers and float value
    s := `gfg 9 true 5.78`

    // Calling NewReader() function for
    // reading each elements of the list
    // and then it place it into "r"
    r := strings.NewReader(s)

    // Declaring different types of variables
    var a string
    var b int
    var c bool
    var d float32

    for {

        // Calling Fscanln() function
        n, err := fmt.Fscanln(r, &a, &b, &c, &d)

        // Checking returned error is
        // end of the line (EOF) or not
        if err == io.EOF {
            break
        }

        // Checking if there is any error
        if err != nil {
            panic(err)
        }

        // Printing the number of items successfully
        // scanned and each elements too
        fmt.Printf("%d: %s, %d, %t, %g", n, a, b, c, d)
    }
}
