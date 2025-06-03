// Go program to illustrate the
// concept of the promoted fields
package main

import "fmt"

// Structure
type details struct {

    // Fields of the
    // details structure
    name   string
    age    int
    gender string
}

// Nested structure
type student struct {
    branch string
    year   int
    details
}

func main() {

    // Initializing the fields of
    // the student structure
    values := student{
        branch: "CSE",
        year:   2010,
        details: details{
        
            name:   "Sumit",
            age:    28,
            gender: "Male",
        },
    }

    // Promoted fields of the student structure
    fmt.Println("Name: ", values.name)
    fmt.Println("Age: ", values.age)
    fmt.Println("Gender: ", values.gender)

    // Normal fields of
    // the student structure
    fmt.Println("Year: ", values.year)
    fmt.Println("Branch : ", values.branch)
}
