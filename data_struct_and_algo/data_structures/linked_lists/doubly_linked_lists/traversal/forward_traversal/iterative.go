package main

import "fmt"

// Node structure in a doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// Prints the list from head to tail
func forwardTraversal(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

// main function to demonstrate traversal
func main() {
	// Create nodes
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}

	// Link nodes
	node1.next = node2

	node2.prev = node1
	node2.next = node3

	node3.prev = node2
	node3.next = node4

	node4.prev = node3

	// Head of the list
	head := node1

	// Perform forward traversal
	fmt.Print("Forward Traversal: ")
	forwardTraversal(head)
}
