package main

import "fmt"

func main() {
	// shorthand declaration of array
	arr := [4]string{"Yoshi", "Mario", "Luigi", "Marwa"}
	fmt.Println("Elements of the array:")

	// accessing the elements of the array using for loop
	for i := 0; i < 3; i++{
		fmt.Println(arr[i])
	}
}
