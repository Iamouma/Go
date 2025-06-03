package main

import "fmt"

func main() {
	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to GeeksforGeeks"
	 
	fmt.Println("String:", mystr)
	 
	    /* if you trying to change
		  the value of the string
		  then the compiler will
		  throw an error, i.e,
		 cannot assign to mystr[1]
	 
	       mystr[1]= 'G'
	       fmt.Println("String:", mystr)
	    */
}
