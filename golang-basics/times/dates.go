// Golang program to compare times
package main

import "fmt"

// importing the time module
import "time"

func main() {
	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)

	// using time.Before() method
	g1 := today.Before(tomorrow)
	fmt.Println("today before tomorrow:", g1)

	// using time.After() method
	g2 := tomorrow.After(today)
	fmt.Println("tomorrow after today:", g2)
}
