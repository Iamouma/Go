// Go progrma to illustrate the use of break statement
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++{
		fmt.Println(i)

		// for loop breaks when the value of i = 3
		if i == 3 {
			break;
		}
	}
}
