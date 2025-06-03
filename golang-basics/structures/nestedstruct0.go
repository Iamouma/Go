// A structure which is the field of aanother is known as Nested Structure(a structure within a structure)
package main

import "fmt"

type Author struct {
	name	string
	branch	string
	year	int
}

type HR struct {
	details Author
}

func main() {
	result := HR{
		details: Author{"Sona", "ECE", 2013},
	}

	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
