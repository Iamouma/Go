package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// function to insert a new node at the beginning
func insertAtBeginning(head *Node, value int) *Node {
	// create a new node
	newNode := &Node{data: value}

	// set next of new node to current head
	newNode.next = head

	// return new node as the new head
	return newNode
}

// Function to print linked list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node = nil

	// Insert nodes at the beginning
	head = insertAtBeginning(head, 10) // 10
	head = insertAtBeginning(head, 20) // 20 -> 10
	head = insertAtBeginning(head, 30) // 30 -> 20 -> 10

	// Print the list
	fmt.Print("Linked List: ")
	printList(head)
}
