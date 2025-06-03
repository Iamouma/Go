// Go program to illustarte the concept of the Pointer to Pointer
package main

import "fmt"

func main() {
	// taking a variable of integer type
	var V int = 100

	// taking a pointer of integer type
	var pt1 *int = &V

	// taking pointer to pointer to pt1 and storing the address of pt1 into pt2
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable V is = ", V)
	fmt.Println("Address of variable V is = ", &V)

	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)

	// deferecencing the pointer to pointer
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)

	// double pointer will give the value of variable V
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}
