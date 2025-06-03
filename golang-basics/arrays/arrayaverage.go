package main

import "fmt"

// function to calculate the average of an array
func calculateAverage(arr [6]int, size int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}
	return sum / size
}

func main() {
	scores := [6]int{67, 59, 29, 35, 4, 34}
	average := calculateAverage(scores, len(scores))
	fmt.Printf("%d\n", average)
}
