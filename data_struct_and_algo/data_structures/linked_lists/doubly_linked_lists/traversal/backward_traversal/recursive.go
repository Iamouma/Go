package main

import "fmt"

// Node defines a doubly linked list node
type Node struct {
	data int
	prev *Node
	next *Node
}

// Recursive function to traverse backward
func backwardTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	backwardTraversal(node.prev)
}

func main() {
	// Create nodes
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}

	// Link the nodes
	node1.next = node2

	node2.prev = node1
	node2.next = node3

	node3.prev = node2
	node3.next = node4

	node4.prev = node3

	// Start recursive backward traversal from the tail
	fmt.Print("Backward Traversal (Recursive): ")
	backwardTraversal(node4)
}

