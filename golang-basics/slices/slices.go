package main

import "fmt"

func main() {
	// slices are like arrays but they are dynamically resizable.
	// slices are a layer of abstraction over arrays.
	// The syntax for defining a slice is
	// i.e var <slice name> []<datatype of data stored in the slice>
	var exampleSlice []int
	// to add elements use the append function
	exampleSlice = append(exampleSlice, 1)
	exampleSlice = append(exampleSlice, 2)
	exampleSlice = append(exampleSlice, 3)

	// also correct i.e exampleSlice := []int{1, 2, 3}
	fmt.Printf("%d is the length of the Slice", len(exampleSlice))              // prints 3
	fmt.Printf("%d is the capacity of the underlying array", cap(exampleSlice))
	fmt.Println(exampleSlice)
}
