// Go program to illustrate how to
// find the capacity of the channel

package main

import "fmt"

// Main function
func main() {

    // Creating a channel
    // Using make() function
    mychnl := make(chan string, 5)
    mychnl <- "Hello"
    mychnl <- "World"
    mychnl <- "Mother"
    mychnl <- "MotherEarth"

    // Finding the capacity of the channel
    // Using cap() function
    fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
