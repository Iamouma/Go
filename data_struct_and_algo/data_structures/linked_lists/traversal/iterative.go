package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Traverse and print the linked list
func traverse(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
}

func main() {
	// Create nodes
	node1 := &Node{data: 5}
	node2 := &Node{data: 10}
	node3 := &Node{data: 15}
	node4 := &Node{data: 20}
	node5 := &Node{data: 25}

	// Link the nodes: 5 -> 10 -> 15 -> 20 -> 25
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	// Traverse and print
	traverse(node1)
}
