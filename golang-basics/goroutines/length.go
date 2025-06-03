// Go program to illustrate how to
// find the length of the channel

package main

import "fmt"
a
// Main function
func main() {

    // Creating a channel
    // Using make() function
    mychnl := make(chan string, 4)
    mychnl <- "Hello"
    mychnl <- "World"
    mychnl <- "Mother"
    mychnl <- "MotherEarth"

    // Finding the length of the channel
    // Using len() function
    fmt.Println("Length of the channel is: ", len(mychnl))
}
