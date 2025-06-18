package main

import "fmt"

// Node structure
type Node struct {
	data int
	prev *Node
	next *Node
}

func findLength(head *Node) int {
	length := 0
	current := head

	for current != nil {
		length++
		current = current.next
	}

	return length
}

func main() {
	// Create nodes
	node1 := &Node{data: 5}
	node2 := &Node{data: 10}
	node3 := &Node{data: 15}

	// Link the nodes
	node1.next = node2

	node2.prev = node1
	node2.next = node3

	node3.prev = node2

	// Find and print the length of the list
	length := findLength(node1)
	fmt.Println("Length of Doubly Linked List:", length)
}
