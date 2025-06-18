package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// DeleteLast deletes the last node from a circular linked list
func DeleteLast(last *Node) *Node {
	if last == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// Only one node
	if last.next == last {
		return nil
	}

	// Traverse to second last node
	curr := last.next
	for curr.next != last {
		curr = curr.next
	}

	// Adjust the next of second last node
	curr.next = last.next
	last = curr

	return last
}

// Helper function to insert at end
func InsertAtEnd(last *Node, data int) *Node {
	newNode := &Node{data: data}
	if last == nil {
		newNode.next = newNode
		return newNode
	}
	newNode.next = last.next
	last.next = newNode
	return newNode
}

// Hel

