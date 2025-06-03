// Golang program to illustrate 
// the use of math.Yn() function 

package main 

import ( 
    "fmt"
    "math"
) 

// Main function 
func main() { 

    // Finding the order-n Bessel 
    // function of the second kind 
    // Using Yn() function 
    res_1 := math.Yn(-3, -2) 
    res_2 := math.Yn(6, 3) 
    res_3 := math.Yn(1, 1.1) 
    res_4 := math.Yn(1, math.NaN()) 
    res_5 := math.Yn(-1, 0) 

    // Displaying the result 
    fmt.Println("Result 1: ", res_1) 
    fmt.Println("Result 2: ", res_2) 
    fmt.Println("Result 3: ", res_3) 
    fmt.Println("Result 4: ", res_4) 
    fmt.Println("Result 5: ", res_5) 

}
