package main

import "fmt"

// Node defines a doubly linked list node
type Node struct {
	data int
	prev *Node
	next *Node
}

// Recursive function to find the length of the doubly linked list
func findLengthRecursive(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + findLengthRecursive(node.next)
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

	// Find and print the length using recursion
	length := findLengthRecursive(node1)
	fmt.Println("Length of Doubly Linked List (Recursive):", length)
}

