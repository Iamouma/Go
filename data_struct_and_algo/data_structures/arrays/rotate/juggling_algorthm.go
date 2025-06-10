package main

import (
	"fmt"
)

// Function to calculate GCD of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Juggling Algorithm for rotating an array to the right by d positions
func rotateRightJuggling(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	d = d % n
	if d == 0 {
		return arr
	}

	// Right rotation by d is equivalent to left rotation by n - d
	k := n - d
	g := gcd(n, k)

	for i := 0; i < g; i++ {
		temp := arr[i]
		j := i

		for {
			next := (j + k) % n
			if next == i {
				break
			}
			arr[j] = arr[next]
			j = next
		}
		arr[j] = temp
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	d := 2

	result := rotateRightJuggling(arr, d)
	fmt.Println("Rotated array:", result) // Output: [6 7 1 2 3 4 5]
}

