// Go program to illustrate 
// How to convert integer 
// variable into String
package main

import (
    "fmt"
)

//Main Function

func main() {
    i := 32
    s := fmt.Sprintf("%d", i)
    fmt.Printf("Type : %T \nValue : %v\n", s, s)

}
