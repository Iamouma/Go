package main
import "fmt"

// Global variable declaration 
var myVariable int = 100 // Global variable

func main() {
    var myVariable int = 200 // Local variable
    fmt.Printf("Local variable takes precedence: %d\n", myVariable) // Accesses local variable
}
