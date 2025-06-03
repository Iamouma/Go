package main

import "fmt"

func main() {
	// string as a range in the for loop
	for index, s := range "WelcomeToMotherEarth" {
		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}
