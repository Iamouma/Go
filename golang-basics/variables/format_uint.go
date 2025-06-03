// Go program to illustrate 
// How to convert integer 
// variable into String
package main

import (
    "fmt"
    "strconv"
)

//Main Function

func main() {
    i := 32
    s := strconv.FormatUint(uint64(i) , 7)
    fmt.Printf("Type : %T \nValue : %v\n", s, s)

}
