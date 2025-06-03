// Go program to illustrate the
// concept of struct equality
// using DeepEqual() method
package main

import (
    "fmt"
    "reflect"
)

// Creating a structure
type Author struct {
    name      string
    branch    string
    language  string
    Particles int
}

// Main function
func main() {

    // Creating variables 
    // of Author structure
    a1 := Author{
        name:      "Soana",
        branch:    "CSE",
        language:  "Perl",
        Particles: 48,
    }

    a2 := Author{
        name:      "Soana",
        branch:    "CSE",
        language:  "Perl",
        Particles: 48,
    }

    a3 := Author{
        name:      "Dia",
        branch:    "CSE",
        language:  "Perl",
        Particles: 48,
    }
    
    // Comparing a1 with a2
    // Using DeepEqual() method
    fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))

    // Comparing a2 with a3
    // Using DeepEqual() method
    fmt.Println("Is a2 equal to a3: ", reflect.DeepEqual(a2, a3))
}
