package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// function to insert a new node at the end
func insertAtEnd(head *Node, value int) *Node {
	// create a new node with next = nil
	newNode := &Node{data: value, next: nil}

	// if the list is empty, new node becomes the head
	if head == nil {
		return newNode
	}

	// traverse to the last node
	current := head
	for current.next != nil {
		current = current.next
	}

	// set the next of the last node to newNode
	current.next = newNode

	return head
}

// Function to print the linked list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node = nil

	// Add nodes to the end of the list
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtEnd(head, 4)
	head = insertAtEnd(head, 5)

	// Insert new node at the end
	head = insertAtEnd(head, 1) // Now: 2 -> 3 -> 4 -> 5 -> 1

	// Print the list
	fmt.Print("Linked List: ")
	printList(head)
}
