package main
import "fmt"

// Basic slice to be used in all examples
var source = []int{10, 20, 30, 40, 50}

func main() {
    destination := make([]int, len(source))

    // Manually copying each element
    for i := 0; i < len(source); i++ {
        destination[i] = source[i]
    }

    fmt.Println("Source:", source)
    fmt.Println("Destination:", destination)
}
