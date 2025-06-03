package main

import "fmt"

func sumAndDiff (a int, b int) (int, int) {
    return a + b, a - b
}

func main() {
    sum, _ := sumAndDiff(5, 3)      // the 2nd variable is ignored and not used

    fmt.Println(sum)                // prints 8
}
