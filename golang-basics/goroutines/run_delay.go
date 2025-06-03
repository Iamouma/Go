package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go display("Goroutine Active")
	display("Main Function")
}
