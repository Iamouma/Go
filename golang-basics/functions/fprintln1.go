// Golang program to illustrate the usage of
// fmt.Fprintln() function

// Including the main package
package main

// Importing fmt and os
import (
    "fmt"
    "os"
)

// Calling main
func main() {

    // Declaring some const variables
    const str1, str2, str3 = "a", "b", "c"

    // Calling Fprintln() function which returns
    // "n" as the number of bytes written and
    // "err" as any error ancountered
    n, err := fmt.Fprintln(os.Stdout, str1, str2, str3)

    // Printing the number of bytes written
    fmt.Print(n, " bytes written.\n")

    // Printing if any error encountered
    fmt.Print(err)

}
