// Golang program to compare times
package main

import "fmt"

// importing time module
import "time"

// Main function
func main() {

    today := time.Now()
    tomorrow := today.Add(24 * time.Hour)
    sameday := tomorrow.Add(-24 * time.Hour)

    if today != tomorrow {
        fmt.Println("today is not tomorrow")
    }

    if sameday == today {
        fmt.Println("sameday is today")
    }

    // using Equal function
    if today.Equal(sameday) {
        fmt.Println("today is sameday")
    }

}
