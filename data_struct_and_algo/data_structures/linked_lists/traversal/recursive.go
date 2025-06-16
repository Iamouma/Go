package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Recursive function to traverse and print the linked list
func traverseRecursive(node *Node) {
	// Base case: if node is nil, stop recursion
	if node == nil {
		return
	}

	// Print current node's data
	fmt.Print(node.data, " ")

	// Recursive call with the next node
	traverseRecursive(node.next)
}

func main() {
	// Create linked list: 5 -> 10 -> 15 -> 20 -> 25
	node1 := &Node{data: 5}
	node2 := &Node{data: 10}
	node3 := &Node{data: 15}
	node4 := &Node{data: 20}
	node5 := &Node{data: 25}

	// Linking nodes
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	// Traverse using recursion
	traverseRecursive(node1)
}
