package main
import "fmt"
  
func main() {
    s1 := "Hello"
    s2 := "Geeks"
    s3 := "Hello"

    // Equality and inequality
    fmt.Println("s1 == s2:", s1 == s2) // false
    fmt.Println("s1 == s3:", s1 == s3) // true
    fmt.Println("s1 != s2:", s1 != s2) // true

    // Lexicographical comparison
    fmt.Println("s1 < s2:", s1 < s2)   // true
    fmt.Println("s1 > s2:", s1 > s2)   // false
    fmt.Println("s1 <= s3:", s1 <= s3) // true
    fmt.Println("s1 >= s3:", s1 >= s3) // true
}
