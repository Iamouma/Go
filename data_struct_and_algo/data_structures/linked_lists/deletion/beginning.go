package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// function to delete the first node
func deleteAtBeginning(head *Node) *Node {
	// chek if list is empty
	if head == nil {
		fmt.Println("List id empty.")
		return nil
	}

	// store current head in temp
	temp := head

	// move head to the next node
	head
}
