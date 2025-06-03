// To access individual fields of a struct you have to use dot(.) operator
package main

import "fmt"

// defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{Name: "Ferrari", Model: "GTC4", Color: "Red", WeightInKg: 1920}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	c.Color = "Black"

	fmt.Println("Car: ", c)
}
