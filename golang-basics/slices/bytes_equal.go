package main
import (
    "bytes"
    "fmt"
)
func main() {
    slc1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slc2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slc3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}

    fmt.Println(bytes.Equal(slc1, slc2)) // Output: true
    fmt.Println(bytes.Equal(slc1, slc3)) // Output: false
}
