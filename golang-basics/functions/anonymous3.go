// You can also pass an anonymous function as an argument to another function
package main

import "fmt"

// Passing anonymous function as an argument
func HWE(i func(p, q string) string) {
	fmt.Println(i("Hello", "World"))
}
func main() {
	value := func(p, q string) string {
		return p + q + "Earth"
	}
	HWE(value)
}
