// Go progrma to illustrate how to close a channel uisng for range loop and close function
package main

import "fmt"

// function
func myfun(mychnl chan string) {
	for v := 0; v < 4; v++ {
		mychnl <- "MotherEarth"
	}
	close(mychnl)
}

// main function
func main() {
	// creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// when the value of ok is set to true means the channel is open and it can
	// send or receive data. When the value of ok is set to false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
