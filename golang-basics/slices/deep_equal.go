package main
import (
    "fmt"
    "reflect"
)
func main() {
    slc1 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slc2 := []byte{'G', 'o', 'l', 'a', 'n', 'g'}
    slc3 := []byte{'g', 'o', 'L', 'A', 'N', 'g'}

    fmt.Println(reflect.DeepEqual(slc1, slc2)) // Output: true
    fmt.Println(reflect.DeepEqual(slc1, slc3)) // Output: false
}
