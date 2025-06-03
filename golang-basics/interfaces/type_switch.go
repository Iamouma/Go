package main

import (
	"fmt"
	"math"
)
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Function to determine the type of shape and calculate area
func calculateArea(shape interface{}) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle area: %.2f\n", s.area())
	case Rectangle:
		fmt.Printf("Rectangle area: %.2f\n", s.area())
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{length: 4, width: 3}

	calculateArea(c)
	calculateArea(r)
}
