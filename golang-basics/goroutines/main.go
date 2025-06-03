package main

import "fmt"

func display(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}

}

func main() {
	go display("Hello, Goroutine!") // Runs concurrently
	display("Hello, Main!")
}
