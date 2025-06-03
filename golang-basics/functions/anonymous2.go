// Anonymous functions can accept arguments
package main

import "fmt"

func main() {
	// Passing arguments in anonymous function
	func(ele string) {
		fmt.Println(ele)
	}("Hello World!")
}
