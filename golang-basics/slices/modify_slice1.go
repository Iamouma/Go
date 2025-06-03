package main
import "fmt"

func addElement(slice []int, element int) []int {
    return append(slice, element)
}

func main() {
    // Initial slice
    a := []int{1, 2, 3}
    
    // Adding an element and returning the new slice
    newNumbers := addElement(a, 4)
    
    fmt.Println("Original slice:", a)
    fmt.Println("New slice:", newNumbers)
}
