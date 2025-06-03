package main
import (
    "fmt"
    "strings"
)

func main() {
    s1 := "Hello, "
    s2 := "Geeks!"

    // Initializing a strings builder
    var builder strings.Builder
    builder.WriteString(s1)
    builder.WriteString(s2)
    fmt.Println("", builder.String())
}
