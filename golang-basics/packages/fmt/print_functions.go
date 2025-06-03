// Go code to visualise the differences
// between various print functions in fmt

package main

import "fmt"

func main() {

    // NOTE:
    // 1. Print(arg) prints arg as
    // it is on the console screen
    fmt.Print("Hello, welcome!")
    // OUTPUT: "Hello, welcome!"
    
    // Some variables to use in our code
    const x, id = "Ashika", 22
    
    // 2. Printf(arg) first formats arg
    // and then prints on the console screen
    fmt.Printf("\n1. Name %q and ID %d found!\n", x, id)
    // OUTPUT:
    // 1. Name "Ashika" and ID 22 found!
    
    // 3. Println(arg) does not format arg
    // It simply prints arg, appends any values
    // mentioned after arg and appends spaces in
    // between and also appends a new line at the end
    fmt.Println("2. Name %q and ID %d found!", x, id)
    // OUTPUT: 2. Name %q and ID %d found! Ashika 22
    
    // Print(arg) does not format, does not add
    // any spaces and does not append any newlines
    fmt.Print("3. Name %q and ID %d found!", x, id)
    // OUTPUT: 3. Name %q and ID %d found!Ashika22
    
    // 4. Sprint(arg) works similar to Print(arg)
    // The key difference is: Sprint returns arg as
    // it is as mentioned in the paranthesis and
    // Print(arg) prints on console screen.
    res := fmt.Sprint("\n4. Name %q and ID %d found!", x, id)
    fmt.Println(res)
    // OUTPUT:
    // 4. Name %q and ID %d found!Ashika22
    
    // 5. Sprintf(arg) works similar to Printf(arg)
    // The key difference is: Sprintf returns arg as
    // a formatted string as mentioned in the paranthesis
    // and Printf prints formatted arg on console screen
    res = fmt.Sprintf("5. Name %q and ID %d found!", x, id)
    fmt.Println(res)
    // OUTPUT: 5. Name "Ashika" and ID 22 found!
    
    // 6. Sprintln(arg) works similar to Println(arg)
    // The key difference is: Sprintln returns arg as
    // it is as mentioned in paranthesis and adds spaces
    // between variables and appends a newline to arg &
    // returns arg value while Println directly prints arg
    // onto the console screen.
    res = fmt.Sprintln("6. Name %q and ID %d found!", x, id)
    fmt.Print(res)
    // OUTPUT: 6. Name %q and ID %d found! Ashika 22
    
    fmt.Print("Goodbye!")
    // OUTPUT: "Goodbye!"
}
