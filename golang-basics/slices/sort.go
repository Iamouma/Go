package main
import (
"fmt"
"sort"
)
func main() {
numbers := []int{5, 2, 6, 3, 1, 4}
sort.Ints(numbers)
fmt.Println(numbers) // Output: [1 2 3 4 5 6]
}
