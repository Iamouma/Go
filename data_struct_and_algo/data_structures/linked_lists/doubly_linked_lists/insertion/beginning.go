package main

import "fmt"

// Node represents a node in a doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// InsertAtBeginning inserts a new node at the beginning of the DLL
func InsertAtBeginning(head *Node, data int) *Node {
	newNode := &Node{data: data}

	newNode.next = head // Point new node's next to current head

	if head != nil {
		head.prev = newNode // Update current head's prev to new node
	}

	newNode.prev = nil // New node is now the first, so prev is nil
	return newNode     // Return new node as the new head
}

// PrintList prints the doubly linked list from head to end
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	head = InsertAtBeginning(head, 30)
	head = InsertAtBeginning(head, 20)
	head = InsertAtBeginning(head, 10)

	fmt.Print("Doubly Linked List: ")
	PrintList(head)
}

