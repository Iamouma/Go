package main

import (
	"fmt"
	"time"
)
func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Task 1 completed"
}

func task2(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "Task completed"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
