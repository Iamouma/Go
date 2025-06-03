// Golang program to illustrate how to 
// find the IEEE 754 binary representation 
package main 

import ( 
    "fmt"
    "math"
) 

// Main function 
func main() { 

    // Finding IEEE 754 binary 
    // representation of the 
    // given numbers 
    // Using Float64bits() function 
    res_1 := math.Float64bits(2) 
    res_2 := math.Float64bits(1) 
    res_3 := math.Float64bits(0) 
    res_4 := math.Float64bits(2.3) 

    // Displaying the result 
    fmt.Println("Result 1: ", res_1) 
    fmt.Println("Result 2: ", res_2) 
    fmt.Println("Result 3: ", res_3) 
    fmt.Println("Result 4: ", res_4) 

}
