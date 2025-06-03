// Go code to demonstrate use & differences
// between various scan functions in fmt package

package main

import "fmt"

// Main function
func main() {

    fmt.Print("Hello, welcome!\n")
    var y, z, age int
    var name string
    
    // Scan function will collect input
    // until a space is encountered
    // from console and stores it in y
    fmt.Scan(&y)
    fmt.Printf("Scan: Y = %d\n", y)
    
    // Scanf collects user input in
    // the specified format and stores it
    // in name and age respectively
    fmt.Scanf("%s", &name)
    fmt.Scanf("%d", &age)
    fmt.Printf("Scanf: Name = %s , Age = %d\n", name, age)
    
    // Scanln will not do nothing if
    // a newline is encountered else
    // it stores input value in z
    fmt.Scanln(&z)
    // In our case after entering age we press
    // enter so this Scanln gets skipped and
    // thus 0 is stored in z as 0 is the
    // default value of int variable in Go
    fmt.Printf("Scanln: Z = %d\n", z)

    var a string
    var g int
    
    // GFG is stored in a
    // 10 is stored in g
    // In case of any error, error message
    // is stored in err1
    // Number of values encountered is stored in res1
    res1, err1 := fmt.Sscan("GFG 10", &a, &g)
    // In case of any error, this block will execute
    if err1 != nil {
        panic(err1)
    }
    fmt.Printf("Sscan -> n = %d , string = %s %d \n", res1, a, g)
    
    // GFG is stored in a
    // 5 is stored in g
    // In case of any error, error message
    // is stored in err2
    // Number of values encountered is stored in res2
    res2, err2 := fmt.Sscanf("String is GFG with 5 iterations", 
                    "String is %s with %d iterations", &a, &g)
    // In case of any error, this block will execute
    if err2 != nil {
        panic(err2)
    }
    fmt.Printf("Sscanf -> n = %d , string = %s %d \n", res2, a, g)
    
    // GFG is stored in a
    // 17 is stored in g
    // In case a newline is encountered,
    // then Sscanln will raise a newline
    // panic error message
    // In case of any error, error message
    // is stored in err3
    // Number of values encountered is stored in res3
    res3, err3 := fmt.Sscanln("GFG 17", &a, &g)
    // In case of any error, this block will execute
    if err3 != nil {
        panic(err3)
    }
    fmt.Printf("Sscanln -> n = %d , string = %s %d \n", res3, a, g)

    fmt.Println("Goodbye!")
}
