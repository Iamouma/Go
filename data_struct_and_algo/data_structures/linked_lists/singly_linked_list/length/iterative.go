package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to find the length of the linked list
func getLength(head *Node) int {
	count := 0
	current := head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func main() {
	// Create linked list: 10 -> 20 -> 30 -> 40 -> 50 -> 60
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}
	node5 := &Node{data: 50}
	node6 := &Node{data: 60}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6

	// Call getLength and print result
	fmt.Println("Length of linked list:", getLength(node1))
}
