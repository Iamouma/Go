// An anonymous function is a function that doesn't have a name.
// It is useful when you wnat to create an inline function.
package main

import "fmt"

func main() {
	// anonymous function
	func() {
		fmt.Println("Hello World!")
	}()
}
