package main
import (
    "bytes"
    "fmt"
)

func main() {
    slice1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slice2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slice3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}

    fmt.Println(bytes.Compare(slice1, slice2)) // Output: 0 (equal)
    fmt.Println(bytes.Compare(slice1, slice3)) // Output: -1 (slice1 < slice3)
}
