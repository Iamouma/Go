package main
import (
	"fmt"
	"sort"
)

type Employee struct {
	Name string
	Age int
}

func main() {
	employees := []Employee	{
		{"John", 28},
		{"Alice", 23},
		{"Bob", 25},
}

sort.Slice(employees, func(i, j int) bool {
return employees[i].Age < employees[j].Age
})

fmt.Println(employees) // Output: [{Alice 23} {Bob 25} {John 28}]
}
