// In Go the main package is a special package which is used with the programs that are executable and this package contains main() function
// The main() function is a special type of function and its is the entry point of the excutable programs. It does not take any argument nor return anything.
// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Main function
func main() {
	// Sorting the given slice
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// Finding the index
	fmt.Println("Index value: ", strings.Index("HelloWorld", "rl"))

	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())
}
