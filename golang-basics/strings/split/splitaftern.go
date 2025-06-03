package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "Welcome,to,GeeksforGeeks"
    fmt.Println("", s)

    result := strings.SplitAfterN(s, ",", 2)
    fmt.Println("Result using SplitAfterN:", result)
}
