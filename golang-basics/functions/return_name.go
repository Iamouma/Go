// Go program to illustrate how to
// give names to the return values
package main

import "fmt"

// myfunc return 2 values of int type
// here, the return value name 
// is rectangle and square
func myfunc(p, q int)( rectangle int, square int ){
    rectangle = p*q
    square = p*p
    return  
}

func main() {
    
    // The return values are assigned into 
    // two different variables
   var area1, area2 = myfunc(2, 4)
   
   // Display the values
   fmt.Printf("Area of the rectangle is: %d", area1 )
   fmt.Printf("\nArea of the square is: %d", area2)
   
}
