// Go program to illustrate how to find the 
// length of the pointer to an array 
package main 
  
import ( 
    "fmt"
) 
  
// Main function 
func main() { 
  
    // Creating and initializing 
    // pointer to array 
    // Using var keyword 
    var ptr1 [6]*int
    var ptr2 [3]*string 
    var ptr3 [4]*float64 
  
    // Finding the length of  
    // the pointer to array 
    // Using len function 
    fmt.Println("Length of ptr1: ", len(ptr1)) 
    fmt.Println("Length of ptr2: ", len(ptr2)) 
    fmt.Println("Length of ptr3: ", len(ptr3)) 
  
} 
