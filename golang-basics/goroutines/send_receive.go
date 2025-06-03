package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("Start Main method")
	// creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}
