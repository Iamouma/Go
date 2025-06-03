package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := -32
	s := strconv.FormatInt(int64(i) , 7)
	fmt.Printf("Type : %T \nValue : %v\n", s, s)
}
