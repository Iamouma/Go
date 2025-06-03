// Golang program to convert a string into Boolean data type
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "MotherEarth"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is: ", str)
	b, _ := strconv.ParseBool(str)
	fmt.Println("After :", reflect.TypeOf(b))
	fmt.Println("Boolean value is: ", b, "\n")

	str1 := "t"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is: ", str1)
	b1, _ := strconv.ParseBool(str1)
	fmt.Println("After :", reflect.TypeOf(b1))
	fmt.Println("Boolen value id: ", b1, "\n")

	str2 := "0"

	fmt.Println("Before :", reflect.TypeOf(str2))
	fmt.Println("String value is: ", str2)
	b2, _ := strconv.ParseBool(str2)
	fmt.Println("After :", reflect.TypeOf(b2))
	fmt.Println("Boolean value is: ", b2, "\n")
}
