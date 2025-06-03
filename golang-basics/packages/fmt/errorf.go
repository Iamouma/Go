// Go code to demonstrate
// the use of errorf function
// in fmt package
package main

import "fmt"

// Main function
func main() {
    const lazy = "I'm very lazy today"
    
    // Errorf formats according to the format specifier and
    // returns the string as a value that satisfies the error
    // and this error string are stored in err
    err := fmt.Errorf("Throwing error because: %q", lazy)

    // err.Error() will return the error message
    // and Println will print it on the output console
    fmt.Println(err.Error())
}
