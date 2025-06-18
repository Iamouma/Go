package main

import "fmt"

// Node structure
type Node struct {
	data int
	prev *Node
	next *Node
}

// Iterative backward traversal function
func backwardTraversal(tail *Node) {
	for tail != nil {
		fmt.Print(tail.data, " ")
		tail = tail.prev
	}
}

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

	// Tail of the list
	tail := node4

	// Perform backward traversal (iterative)
	fmt.Print("Backward Traversal (Iterative): ")
	backwardTraversal(tail)
}
