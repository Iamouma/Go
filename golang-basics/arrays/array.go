package main

import "fmt"

func main() {
	// Arrays are essentially staorage spaces that can be filled with as much data as one would like.
	// An array in golang has fixed size similar to C and C++ that we give while we define it.
	// Arrays in golang are defines using the syntax below
	// var <name of the array> [<size of the array>]<type of data stored in the array>
	// An array named favNums filled with 3 integers
	var favNums [3]int

	// Insert data into the array
	// The first storage space will be assigned the value of 1. It has an index of 0.
	favNums[0] = 1
	// The second storage space will be assigned the value of 2. It has an index of 1.
	favNums[1] = 2
	// The third and final storage space will be assigned the value of 3. It has an index of 2.
	favNums[2] = 3
	fmt.Println(favNums[0])
}
