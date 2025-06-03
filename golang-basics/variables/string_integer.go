package main

import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	str := "MotherEarth"

	fmt.Println("Before :", reflect.TypeOf(str))
	fmt.Println("String value is : ", str)
	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println("After :", reflect.TypeOf(i))
	fmt.Println("Integer value is: ", i, "\n")

	str1 := "100"

	fmt.Println("Before :", reflect.TypeOf(str1))
	fmt.Println("String value is ", str1)
	i1, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println("After :", reflect.TypeOf(i1))
	fmt.Println("Integer value is: ", i1, "\n")
}
