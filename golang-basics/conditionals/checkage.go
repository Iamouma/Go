package main

import "fmt"

func main() {
	userAge := 26

	if userAge < 20 {
	    fmt.Println("Below 20")
	} else if userAge >= 20 && userAge <= 60 {
	    fmt.Println("Between 20 and 60")
	} else {
	    fmt.Println("Above 60")
	}
}
