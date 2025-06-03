package main

import "fmt"

func main() {
	// creating an array whose size is represented by the ellipsis
	myarray := [...]int{29, 79, 49, 39, 20, 49, 48, 49}

	// iterate array using for loop
	for x := 0; x < len(myarray); x++ {
		fmt.Printf("%d\n", myarray[x])
	}
}
