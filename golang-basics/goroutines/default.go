package main
import (
	"fmt"
)
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("No tasks are ready yet")
	}
}
