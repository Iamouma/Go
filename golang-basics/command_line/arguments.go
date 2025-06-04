// Command-line arguments are a common way to parameterize execution of programs.
// os.Args provide access to raw command-line arguments.
// Note that the first value in this slice is the path to the program and os.Args[1:] holds the arguments to the program.


package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]


	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
