package main

import "fmt"

func main() {
	// Golang also provides a for-range loop to loop over an array, slice or few other datastructures.

	// The for-range loop provides us access to the index and value of the elements in the array or slice we are looping through as shown below/
	myList := []int{1,2,3}

    	for index, value := range myList {
        	fmt.Printf("%d is index, %d is value", index, value)
    	}

	// We can ignore one or both the fields in a for-range loop by giving an _ instead of giving a variable name
	for _, value := range myList {
		fmt.Println(value)
	}
}

