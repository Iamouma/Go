package main

import "fmt"

// Node strcuture
type Node struct {
	data int
	next *Node
}

// Recurse function to count the node
func countNodes(node *Node) int {
	// Base case: if node is nil, return 0
	if node == nil {
		return 0
	}

	// recursive case: 1 + count of remaining nodes
	return 1 + countNodes(node.next)
}

func main() {
	// Create linked list: 1 -> 3 -> 1 -> 2 -> 1
	node1 := &Node{data: 1}
	node2 := &Node{data: 3}
	node3 := &Node{data: 1}
	node4 := &Node{data: 2}
	node5 := &Node{data: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	// Call countNodes and print result
	fmt.Println("Length of linked list:", countNodes(node1))
}
